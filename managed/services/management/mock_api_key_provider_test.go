// Code generated by mockery v2.36.0. DO NOT EDIT.

package management

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockApiKeyProvider is an autogenerated mock type for the apiKeyProvider type
type mockApiKeyProvider struct {
	mock.Mock
}

// CreateAdminAPIKey provides a mock function with given fields: ctx, name, ttl
func (_m *mockApiKeyProvider) CreateAdminAPIKey(ctx context.Context, name string, ttl time.Duration) (int64, string, error) {
	ret := _m.Called(ctx, name, ttl)

	var r0 int64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) (int64, string, error)); ok {
		return rf(ctx, name, ttl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) int64); ok {
		r0 = rf(ctx, name, ttl)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) string); ok {
		r1 = rf(ctx, name, ttl)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, time.Duration) error); ok {
		r2 = rf(ctx, name, ttl)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// newMockApiKeyProvider creates a new instance of mockApiKeyProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockApiKeyProvider(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockApiKeyProvider {
	mock := &mockApiKeyProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
