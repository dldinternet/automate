// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/legacy/legacy.proto

package legacy

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.legacy.LegacyDataCollector/Status", "ingest:status", "read", "GET", "/events/data-collector", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}