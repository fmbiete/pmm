// Code generated by mockery v2.33.0. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAgentService is an autogenerated mock type for the agentService type
type mockAgentService struct {
	mock.Mock
}

// Logs provides a mock function with given fields: ctx, pmmAgentID, agentID, limit
func (_m *mockAgentService) Logs(ctx context.Context, pmmAgentID string, agentID string, limit uint32) ([]string, uint32, error) {
	ret := _m.Called(ctx, pmmAgentID, agentID, limit)

	var r0 []string
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint32) ([]string, uint32, error)); ok {
		return rf(ctx, pmmAgentID, agentID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint32) []string); ok {
		r0 = rf(ctx, pmmAgentID, agentID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, uint32) uint32); ok {
		r1 = rf(ctx, pmmAgentID, agentID, limit)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, uint32) error); ok {
		r2 = rf(ctx, pmmAgentID, agentID, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// newMockAgentService creates a new instance of mockAgentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAgentService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAgentService {
	mock := &mockAgentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
