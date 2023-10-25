// Code generated by mockery v2.36.0. DO NOT EDIT.

package backup

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"

	servicesbackup "github.com/percona/pmm/managed/services/backup"
)

// mockBackupService is an autogenerated mock type for the backupService type
type mockBackupService struct {
	mock.Mock
}

// PerformBackup provides a mock function with given fields: ctx, params
func (_m *mockBackupService) PerformBackup(ctx context.Context, params servicesbackup.PerformBackupParams) (string, error) {
	ret := _m.Called(ctx, params)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, servicesbackup.PerformBackupParams) (string, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, servicesbackup.PerformBackupParams) string); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, servicesbackup.PerformBackupParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreBackup provides a mock function with given fields: ctx, serviceID, artifactID, pitrTimestamp
func (_m *mockBackupService) RestoreBackup(ctx context.Context, serviceID string, artifactID string, pitrTimestamp time.Time) (string, error) {
	ret := _m.Called(ctx, serviceID, artifactID, pitrTimestamp)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time) (string, error)); ok {
		return rf(ctx, serviceID, artifactID, pitrTimestamp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time) string); ok {
		r0 = rf(ctx, serviceID, artifactID, pitrTimestamp)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, time.Time) error); ok {
		r1 = rf(ctx, serviceID, artifactID, pitrTimestamp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SwitchMongoPITR provides a mock function with given fields: ctx, serviceID, enabled
func (_m *mockBackupService) SwitchMongoPITR(ctx context.Context, serviceID string, enabled bool) error {
	ret := _m.Called(ctx, serviceID, enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, serviceID, enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockBackupService creates a new instance of mockBackupService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockBackupService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockBackupService {
	mock := &mockBackupService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
