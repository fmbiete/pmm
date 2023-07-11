// Code generated by mockery v2.31.1. DO NOT EDIT.

package management

import (
	context "context"
	time "time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
	mock "github.com/stretchr/testify/mock"
)

// mockVictoriaMetricsClient is an autogenerated mock type for the victoriaMetricsClient type
type mockVictoriaMetricsClient struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, query, ts, opts
func (_m *mockVictoriaMetricsClient) Query(ctx context.Context, query string, ts time.Time, opts ...v1.Option) (model.Value, v1.Warnings, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query, ts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 model.Value
	var r1 v1.Warnings
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, ...v1.Option) (model.Value, v1.Warnings, error)); ok {
		return rf(ctx, query, ts, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, ...v1.Option) model.Value); ok {
		r0 = rf(ctx, query, ts, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, ...v1.Option) v1.Warnings); ok {
		r1 = rf(ctx, query, ts, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, time.Time, ...v1.Option) error); ok {
		r2 = rf(ctx, query, ts, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// newMockVictoriaMetricsClient creates a new instance of mockVictoriaMetricsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockVictoriaMetricsClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockVictoriaMetricsClient {
	mock := &mockVictoriaMetricsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
