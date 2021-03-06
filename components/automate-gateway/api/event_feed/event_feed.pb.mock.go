// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/event_feed/event_feed.proto

package event_feed

import (
	"context"

	request "github.com/chef/automate/components/automate-gateway/api/event_feed/request"
	response "github.com/chef/automate/components/automate-gateway/api/event_feed/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the EventFeedServer interface (at compile time)
var _ EventFeedServer = &EventFeedServerMock{}

// NewEventFeedServerMock gives you a fresh instance of EventFeedServerMock.
func NewEventFeedServerMock() *EventFeedServerMock {
	return &EventFeedServerMock{validateRequests: true}
}

// NewEventFeedServerMockWithoutValidation gives you a fresh instance of
// EventFeedServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewEventFeedServerMockWithoutValidation() *EventFeedServerMock {
	return &EventFeedServerMock{}
}

// EventFeedServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type EventFeedServerMock struct {
	validateRequests          bool
	GetEventFeedFunc          func(context.Context, *request.EventFilter) (*response.Events, error)
	GetEventTypeCountsFunc    func(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	GetEventTaskCountsFunc    func(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	GetEventStringBucketsFunc func(context.Context, *request.EventStrings) (*response.EventStrings, error)
}

func (m *EventFeedServerMock) GetEventFeed(ctx context.Context, req *request.EventFilter) (*response.Events, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetEventFeedFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetEventFeed' not implemented")
}

func (m *EventFeedServerMock) GetEventTypeCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetEventTypeCountsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetEventTypeCounts' not implemented")
}

func (m *EventFeedServerMock) GetEventTaskCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetEventTaskCountsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetEventTaskCounts' not implemented")
}

func (m *EventFeedServerMock) GetEventStringBuckets(ctx context.Context, req *request.EventStrings) (*response.EventStrings, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetEventStringBucketsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetEventStringBuckets' not implemented")
}

// Reset resets all overridden functions
func (m *EventFeedServerMock) Reset() {
	m.GetEventFeedFunc = nil
	m.GetEventTypeCountsFunc = nil
	m.GetEventTaskCountsFunc = nil
	m.GetEventStringBucketsFunc = nil
}
