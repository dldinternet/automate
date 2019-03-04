package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
	go_oidc "github.com/coreos/go-oidc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"

	"github.com/chef/automate/lib/logger"
)

// toggle for logs or no logs in test output
const wantLogs = false

const bldrURLString = "https://builder.test/"

func TestSafeCmpStrings(t *testing.T) {
	t.Run("when the strings are the same, it returns true", func(t *testing.T) {
		require.Equal(t, safeCmpStrings("teststring", "teststring"), true)
	})

	t.Run("when the strings are the same length but have different content, returns false", func(t *testing.T) {
		require.Equal(t, safeCmpStrings("length7", "7length"), false)
	})

	t.Run("when the strings are the different lengths, returns false", func(t *testing.T) {
		require.Equal(t, safeCmpStrings("somestring", "wrong"), false)
	})

	t.Run("when one string is empty, returns false", func(t *testing.T) {
		require.Equal(t, safeCmpStrings("", "wrong"), false)
	})

	t.Run("when both strings are empty, returns true", func(t *testing.T) {
		require.Equal(t, safeCmpStrings("", ""), true)
	})
}

func TestHealthHandler(t *testing.T) {
	ms := memstore.New(time.Minute)
	hdlrFunc := newTestServer(t, ms).mux.ServeHTTP

	t.Run("GET /health", func(t *testing.T) {
		expected := `{"status":"SERVING"}`
		require.HTTPSuccess(t, hdlrFunc, "GET", "/health", nil)
		require.HTTPBodyContains(t, hdlrFunc, "GET", "/health", nil, expected)
	})
}

func TestNewHandler(t *testing.T) {
	ms := memstore.New(time.Minute)
	s := newTestServer(t, ms)
	hdlr := s.mux

	cases := map[string]func(*testing.T){
		"GET /new?state=xyz is redirected to dex, stores generated relay_state xyz": func(t *testing.T) {
			r := httptest.NewRequest("GET", "/new?state=%2Fnodes", nil) // = "/nodes"
			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()
			require.Equal(t, http.StatusSeeOther, resp.StatusCode, "is redirected")
			require.Contains(t, resp.Header.Get("Location"), "dex/auth", "to dex auth endpoint")

			var sessionID string
			for _, k := range resp.Cookies() {
				if k.Name == "session" {
					sessionID = k.Value
				}
			}
			require.NotEmpty(t, sessionID, "there is a session ID cookie")
			data, exists, err := ms.Find(sessionID)
			require.Nil(t, err)
			require.True(t, exists, "there is a stored session")
			var cookie struct {
				Data struct {
					RelayState  string `json:"relay_state"`
					ClientState string `json:"client_state"`
				} `json:"data"`
			}
			err = json.Unmarshal(data, &cookie)
			require.Nil(t, err)
			require.NotEmpty(t, cookie.Data.RelayState)
			require.Contains(t, resp.Header.Get("Location"), cookie.Data.RelayState)
			require.Equal(t, "/nodes", cookie.Data.ClientState)
		},
		"GET /new from builder is redirected to dex, stores generated relay_state xyz": func(t *testing.T) {
			reqStr := fmt.Sprintf("/new?state=xyz&client_id=bldr-client&redirect_uri=%s&response_type=code&scope=openid&nonce=0",
				bldrURLString)
			r := httptest.NewRequest("GET", reqStr, nil)
			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusSeeOther, resp.StatusCode, "is redirected")
			require.Contains(t, resp.Header.Get("Location"), "dex/auth", "to dex auth endpoint")

			var sessionID string
			for _, k := range resp.Cookies() {
				if k.Name == "session" {
					sessionID = k.Value
				}
			}
			require.NotEmpty(t, sessionID, "there is a session ID cookie")
			data, exists, err := ms.Find(sessionID)
			require.NoError(t, err)
			require.True(t, exists, "there is a stored session")
			var cookie struct {
				Data struct {
					RelayState  string `json:"relay_state"`
					ClientState string `json:"client_state"`
					RedirectURI string `json:"redirect_uri"`
				} `json:"data"`
			}
			err = json.Unmarshal(data, &cookie)
			require.Nil(t, err)
			require.NotEmpty(t, cookie.Data.RelayState)
			require.Contains(t, resp.Header.Get("Location"), cookie.Data.RelayState)
			require.Equal(t, "xyz", cookie.Data.ClientState)
			require.Equal(t, bldrURLString, cookie.Data.RedirectURI)
		},

		"GET /new from builder fails with a 401 when client_id matches bldr but redirect_uri does not match bldr config": func(t *testing.T) {
			reqStr := fmt.Sprintf("/new?state=xyz&client_id=bldr-client&redirect_uri=%s&response_type=code&scope=openid&nonce=0",
				"https://wrong/")
			r := httptest.NewRequest("GET", reqStr, nil)
			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()
			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		"GET /new from builder fails with a 401 when client_id is passed but does not match bldr": func(t *testing.T) {
			reqStr := fmt.Sprintf("/new?state=xyz&client_id=not-bldr-client&redirect_uri=%s&response_type=code&scope=openid&nonce=0",
				bldrURLString)
			r := httptest.NewRequest("GET", reqStr, nil)
			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()
			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		// This test is quite boring, also, we don't test DELETE, PUT, ... -- however
		// it also doesn't hurt.
		"POST /new is forbidden": func(t *testing.T) {
			r := httptest.NewRequest("POST", "/new", nil)
			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()
			require.Equal(t, http.StatusUnauthorized, resp.StatusCode, "is forbidden")
		},

		"GET /new deletes the previous session": func(t *testing.T) {
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			relayState := "relaaay"
			clientState := "/nodes"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("GET", "/new", nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

			w := httptest.NewRecorder()

			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			hasNoCookieWithSessionID(t, resp, sessionID)

			_, exists, err := ms.Find(sessionID)
			require.NoError(t, err)
			require.False(t, exists)
		},
	}

	for name, tc := range cases {
		tc := tc // capture range variable (see https://blog.golang.org/subtests)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tc(t)
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	incomingIDToken := "ey.xyz"
	newRefreshRequest := func() *http.Request {
		r := httptest.NewRequest("GET", "/refresh", nil)
		r.Header.Set("Authorization", fmt.Sprintf("bearer %s", incomingIDToken))
		return r
	}

	ms := memstore.New(time.Minute)
	s := newTestServer(t, ms)
	hdlr := s.mux

	cases := map[string]func(*testing.T){
		"GET /refresh without a session cookie is forbidden": func(t *testing.T) {
			r := newRefreshRequest()
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},
		"GET /refresh with an unknown session cookie is forbidden": func(t *testing.T) {
			r := newRefreshRequest()
			r.AddCookie(&http.Cookie{Name: "session", Value: "nicetry"})
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},

		"GET /refresh with a good session cookie": func(t *testing.T) {
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			refreshToken := "somethingopaque"
			newRefreshToken := "somethingopaqueandnewforrefresh"
			newIDToken := "somethingopaqueandnew" // it's a JWT, but we don't care here

			t1 := oauth2.Token{RefreshToken: newRefreshToken}
			t2 := t1.WithExtra(map[string]interface{}{"id_token": newIDToken})
			s.setTestToken(t2, nil)

			sessionData := struct {
				RT string `json:"refresh_token"` // should equal refreshTokenKey const
			}{refreshToken}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}
			r := newRefreshRequest()
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"id_token": %q}`, newIDToken), string(body))
			hasNoCookieWithSessionID(t, resp, sessionID)
			hasCookieWithAnySessionID(t, resp)
			hasCookieWithSecureSessionSettings(t, resp)
		},

		// Note: the situation "good session cookie + no refresh_token + expired id_token"
		//       is covered by the "token exchange failing" test case below (since that
		//       is how it will look like).
		"GET /refresh with a good session cookie that doesn't have a refresh_token": func(t *testing.T) {
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-50Q"

			// set up the oauth2 token source to return the existing id_token (mocking
			// the situation where the id_token hasn't expired yet)
			t1 := (&oauth2.Token{}).WithExtra(map[string]interface{}{"id_token": incomingIDToken})
			s.setTestToken(t1, nil)

			// anything, just so it doesn't trigger "no session found" (which happens if
			// there's no keys at all in session data)
			sessionData := struct {
				Foo string `json:"foo"`
			}{"bar"}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}
			r := newRefreshRequest()
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"id_token": %q}`, incomingIDToken), string(body))
			hasNoCookieWithSessionID(t, resp, sessionID)
			hasCookieWithAnySessionID(t, resp)
			hasCookieWithSecureSessionSettings(t, resp)
		},

		"GET /refresh with a good session cookie, but token exchange failing": func(t *testing.T) {
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			refreshToken := "somethingopaque"

			s.setTestToken(nil, errors.New("token exchange failed"))

			sessionData := struct {
				RT string `json:"refresh_token"` // should equal refreshTokenKey const
			}{refreshToken}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}
			r := newRefreshRequest()
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			// httptest's http.Response.Body can always be read safely
			hasNoCookieWithSessionID(t, resp, sessionID)
			hasCookieWithAnySessionID(t, resp)
			hasCookieWithSecureSessionSettings(t, resp)
		},
	}

	for name, tc := range cases {
		t.Run(name, tc)
	}
}

func TestCallbackHandler(t *testing.T) {
	ms := memstore.New(time.Minute)
	s := newTestServer(t, ms)
	hdlr := s.mux

	cases := map[string]func(*testing.T){

		"GET /callback without a session cookie is forbidden": func(t *testing.T) {
			r := httptest.NewRequest("GET", "/callback", nil)
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},

		"GET /callback with a session cookie but missing code and state is rejected": func(t *testing.T) {
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			r := httptest.NewRequest("GET", "/callback", nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},

		`GET /callback?code=X&state=Y with a session cookie when code exchange
     succeeds, stores refresh_token and redirects to signin`: func(t *testing.T) {
			code := "randomstringcode"
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			relayState := "relaaay"
			clientState := "/nodes"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			newRefreshToken := "somethingopaqueandnewforrefresh"
			newIDToken := "somethingopaqueandnew" // it's a JWT, but we don't care here
			t1 := oauth2.Token{RefreshToken: newRefreshToken}
			t2 := t1.WithExtra(map[string]interface{}{"id_token": newIDToken})
			s.setTestToken(t2, nil)

			r := httptest.NewRequest("GET", fmt.Sprintf("/callback?code=%s&state=%s", code, relayState), nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusSeeOther, resp.StatusCode)
			require.Contains(t, resp.Header.Get("Location"), fmt.Sprintf("/signin#id_token=%s&state=%s", newIDToken, clientState))

			// this process gives us a new session id, so we just check that the old one
			// is no more, and SOME other one is set
			hasNoCookieWithSessionID(t, resp, sessionID)
			hasCookieWithAnySessionID(t, resp)
			hasCookieWithSecureSessionSettings(t, resp)
		},

		`GET /callback?code=X&state=Y with a session cookie
     when relay state mismatches, stores nothing`: func(t *testing.T) {
			code := "randomstringcode"
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			storedRelayState := "relaaay"
			givenRelayState := "reeelay"
			clientState := "/nodes"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{storedRelayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			s.setTestToken(nil, errors.New("code exchange failed"))

			r := httptest.NewRequest("GET", fmt.Sprintf("/callback?code=%s&state=%s", code, givenRelayState), nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},

		`GET /callback?code=X&state=Y with a session cookie
     when code exchange fails, stores nothing`: func(t *testing.T) {
			code := "randomstringcode"
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			relayState := "relaaay"
			clientState := "/nodes"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			s.setTestToken(nil, errors.New("code exchange failed"))

			r := httptest.NewRequest("GET", fmt.Sprintf("/callback?code=%s&state=%s", code, relayState), nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			require.Empty(t, resp.Cookies(), "no cookie was given")
		},

		`GET /callback?code=X&state=Y, referred by builder, stores code in tokenCache & redirects with state and code`: func(t *testing.T) {
			code := "randomstringcode"
			sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
			relayState := "relaaay"
			clientState := "/nodes"
			redirectURI := "builder.test"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
				RU string `json:"redirect_uri"`
			}{relayState, clientState, redirectURI}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			newIDToken := "newIdTokenTest"
			t1 := oauth2.Token{}
			t2 := t1.WithExtra(map[string]interface{}{"id_token": newIDToken})
			s.setTestToken(t2, nil)
			s.setTokenInMap(code, t2, nil)

			r := httptest.NewRequest("GET", fmt.Sprintf("/callback?code=%s&state=%s", code, relayState), nil)
			r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			token, found := s.tokenCache.Get(code)
			require.True(t, found)
			require.Equal(t, token, t2)

			require.Equal(t, http.StatusSeeOther, resp.StatusCode)
			require.Contains(t, resp.Header.Get("Location"), redirectURI)

			hasNoCookieWithSessionID(t, resp, sessionID)
			hasCookieRemovingAnySessionIDs(t, resp)
			hasCookieWithSecureSessionSettings(t, resp)

			_, exists, err := ms.Find(sessionID)
			require.NoError(t, err)
			require.False(t, exists)
		},
	}

	for name, tc := range cases {
		t.Run(name, tc)
	}
}

func TestTokenHandler(t *testing.T) {
	ms := memstore.New(time.Minute)
	s := newTestServer(t, ms)
	relayState := "relaaay"
	clientState := "xyz"
	hdlr := s.mux
	sessionID := "RBUh6l2c2JB3h6gEWAOGt2HHtL4inzSIgk-oNB-51Q"
	newTokenRequest := func(code string) *http.Request {
		data := url.Values{}
		data.Set("code", code)

		reqBody := strings.NewReader(data.Encode())
		reqURL, err := url.Parse("/token")
		require.NoError(t, err)
		q := reqURL.Query()
		q.Add("redirect_uri", "https://builder.test/")
		q.Add("client_id", "bldr-client")
		q.Add("client_secret", "bldr-secret")
		reqURL.RawQuery = q.Encode()
		r := httptest.NewRequest("POST", reqURL.String(), reqBody)
		r.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		return r
	}

	cases := map[string]func(*testing.T){

		`POST /token with code that matches an access token that's stored in tokenCache`: func(t *testing.T) {
			code := "testing"
			newIDToken := "ey.xyz"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
				AT string `json:"access_token"`
			}{relayState, clientState, newIDToken}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			t1 := oauth2.Token{}
			t2 := t1.WithExtra(map[string]interface{}{"id_token": newIDToken})
			s.setTestToken(t2, nil)
			s.setTokenInMap(code, t2, nil)

			r := newTokenRequest(code)

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"access_token": %q}`, newIDToken), string(body))
		},

		`POST /token with code that doesn't match tokenCache errors out`: func(t *testing.T) {
			code := "testing2"
			newIDToken := "ey.xyz"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			t1 := oauth2.Token{}
			t2 := t1.WithExtra(map[string]interface{}{"id_token": newIDToken})
			s.setTestToken(t2, nil)

			r := newTokenRequest(code)

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, "failed to get token\n", string(body))
		},

		`POST /token with an access token that has no id token errors out`: func(t *testing.T) {
			code := "testing3"
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			t1 := &oauth2.Token{}
			s.setTestToken(t1, nil)
			s.setTokenInMap(code, t1, nil)

			r := newTokenRequest(code)

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, "no id_token in token response\n", string(body))
		},

		`POST /token with no code errors out`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		},

		`POST /token when bldr is not configured returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}
			s.bldrClient = nil

			r := newTokenRequest("")

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

			s.bldrClient = getBldrStruct(t)
		},

		`POST /token when client_id is not passed returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Del("client_id")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		`POST /token when client_id that does not match the bldr client id returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Set("client_id", "wrong-id")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		`POST /token when redirect_uri is not passed returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Del("redirect_uri")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		`POST /token when redirect_id does not match bldr config returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Set("redirect_uri", "https://wrong.org/")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		`POST /token when client_secret is missing returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Del("client_secret")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},

		`POST /token when client_secret is wrong returns a 401`: func(t *testing.T) {
			sessionData := struct {
				RS string `json:"relay_state"`
				CS string `json:"client_state"`
			}{relayState, clientState}
			if err := addSessionDataToStore(ms, sessionID, sessionData, nil); err != nil {
				t.Fatal(err)
			}

			r := newTokenRequest("")
			reqURL := r.URL
			q := reqURL.Query()
			q.Set("client_secret", "this is wrong")
			reqURL.RawQuery = q.Encode()
			r.URL = reqURL

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},
	}

	for name, tc := range cases {
		t.Run(name, tc)
	}
}

func TestUserInfoHandler(t *testing.T) {
	ms := memstore.New(time.Minute)
	s := newTestServer(t, ms)
	hdlr := s.mux

	cases := map[string]func(*testing.T){

		`GET /userinfo with unverified id token returns unauthorized error`: func(t *testing.T) {
			r := httptest.NewRequest("GET", "/userinfo", nil)
			r.Header.Set("Authorization", fmt.Sprintf("bearer %s", "unverifiedToken"))

			s.setTestIDToken(nil, errors.New("unverified token"))

			w := httptest.NewRecorder()
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		},
	}

	for name, tc := range cases {
		t.Run(name, tc)
	}
}

// helpers
func hasNoCookieWithSessionID(t *testing.T, resp *http.Response, unExpectedSessionID string) {
	var found bool
	t.Helper()
	for _, k := range resp.Cookies() {
		if k.Name == "session" && k.Value == unExpectedSessionID {
			found = true
		}
	}
	require.Falsef(t, found, "session ID %s should not have been there", unExpectedSessionID)
}

func hasCookieWithAnySessionID(t *testing.T, resp *http.Response) {
	var sessionID string
	t.Helper()
	for _, k := range resp.Cookies() {
		if k.Name == "session" {
			sessionID = k.Value
		}
	}
	require.NotEmpty(t, sessionID, "there is a session ID cookie")
}

func hasCookieRemovingAnySessionIDs(t *testing.T, resp *http.Response) {
	assertion := false
	t.Helper()
	for _, k := range resp.Cookies() {
		if k.Name == "session" {
			if k.Value == "" {
				assertion = true
			}
		}
	}
	require.True(t, assertion, "there is a `session=` cookie (removing session cookie)")
}

func hasCookieWithSecureSessionSettings(t *testing.T, resp *http.Response) {
	t.Helper()
	for _, k := range resp.Cookies() {
		if k.Name == "session" {
			require.True(t, k.Secure)
			require.True(t, k.HttpOnly)
			// if we set Persist: true on the scsManager,
			// the cookie will have the following keys.
			// Currently, these default to 24h in SCS
			// https://godoc.org/github.com/alexedwards/scs#Manager.Persist
			// until we modify those defaults, checking that they're not
			// empty is enough
			require.NotEmpty(t, k.MaxAge)
			require.NotEmpty(t, k.Expires)
		}
	}
}

// Note: it's quite unfortunate how this reaches into the internals of scs; but
// I haven't found a better way...
func addSessionDataToStore(store scs.Store, sessionID string, sessionData interface{}, t *time.Time) error {
	if t == nil {
		t0 := time.Now()
		t = &t0 // t = &time.Now() doesn't work
	}
	data := struct {
		Data     interface{} `json:"data"`
		Deadline int64       `json:"deadline"`
	}{Data: sessionData, Deadline: t.Add(time.Minute).UnixNano()}
	bs, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	return errors.Wrap(store.Save(sessionID, bs, t.Add(time.Minute)), "store session")
}

func newTestServer(t *testing.T, store scs.Store) *Server {
	var l logger.Logger
	var err error
	if wantLogs {
		l, err = logger.NewLogger("text", "debug")
		if err != nil {
			t.Fatal(err)
		}
	} else {
		l = logger.NewTestLogger()
	}

	u, err := url.Parse("/signin")
	require.NoError(t, err)

	// This isn't an ideal test setup as we're calling the helper
	// in code & test, but it lets us assert that the helper is doing
	// the right thing
	scsManager := createSCSManager(store)

	s := Server{
		log:        l,
		mgr:        scsManager,
		signInURL:  u,
		bldrClient: getBldrStruct(t),
		client:     &testOAuth2Config{token: &oauth2.Token{}},
		tokenCache: cache.New(1*time.Minute, 5*time.Minute),
	}
	s.initHandlers()

	return &s
}

func getBldrStruct(t *testing.T) *BldrClient {
	bURL, err := url.Parse(bldrURLString)
	require.NoError(t, err)

	return &BldrClient{
		SignInURL:    bURL,
		ClientID:     "bldr-client",
		ClientSecret: "bldr-secret",
	}
}

func (s *Server) setTestToken(token *oauth2.Token, err error) {
	s.client = &testOAuth2Config{token: token, err: err, idToken: &go_oidc.IDToken{}}
}

func (s *Server) setTokenInMap(code string, token *oauth2.Token, err error) {
	s.tokenCache.Set(code, token, cache.DefaultExpiration)
}

func (s *Server) setTestIDToken(idToken *go_oidc.IDToken, err error) {
	s.client = &testOAuth2Config{idToken: idToken, err: err}
}

type testOAuth2Config struct {
	token   *oauth2.Token
	err     error
	idToken *go_oidc.IDToken
}

func (x *testOAuth2Config) TokenSource(context.Context, *oauth2.Token) oauth2.TokenSource {
	return x
}

func (x *testOAuth2Config) Exchange(context.Context, string) (*oauth2.Token, error) {
	return x.token, x.err
}

func (x *testOAuth2Config) Token() (*oauth2.Token, error) {
	return x.token, x.err
}

func (x *testOAuth2Config) Client() *http.Client {
	return http.DefaultClient
}

func (x *testOAuth2Config) AuthCodeURL(state string, _ ...oauth2.AuthCodeOption) string {
	return fmt.Sprintf("https://localhost/dex/auth?blahblah&state=%s", state)
}

func (x *testOAuth2Config) Verify(context.Context, string) (*go_oidc.IDToken, error) {
	return x.idToken, x.err
}