// Code generated by mockery v2.32.0. DO NOT EDIT.

package checks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
)

// mockAgentsRegistry is an autogenerated mock type for the agentsRegistry type
type mockAgentsRegistry struct {
	mock.Mock
}

// StartMongoDBQueryBuildInfoAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, files, tdp
func (_m *mockAgentsRegistry) StartMongoDBQueryBuildInfoAction(ctx context.Context, id string, pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, files, tdp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, *models.DelimiterPair) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, files, tdp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMongoDBQueryGetCmdLineOptsAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, files, tdp
func (_m *mockAgentsRegistry) StartMongoDBQueryGetCmdLineOptsAction(ctx context.Context, id string, pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, files, tdp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, *models.DelimiterPair) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, files, tdp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMongoDBQueryGetDiagnosticDataAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, files, tdp
func (_m *mockAgentsRegistry) StartMongoDBQueryGetDiagnosticDataAction(ctx context.Context, id string, pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, files, tdp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, *models.DelimiterPair) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, files, tdp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMongoDBQueryGetParameterAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, files, tdp
func (_m *mockAgentsRegistry) StartMongoDBQueryGetParameterAction(ctx context.Context, id string, pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, files, tdp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, *models.DelimiterPair) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, files, tdp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMongoDBQueryReplSetGetStatusAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, files, tdp
func (_m *mockAgentsRegistry) StartMongoDBQueryReplSetGetStatusAction(ctx context.Context, id string, pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, files, tdp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, *models.DelimiterPair) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, files, tdp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMySQLQuerySelectAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify
func (_m *mockAgentsRegistry) StartMySQLQuerySelectAction(ctx context.Context, id string, pmmAgentID string, dsn string, query string, files map[string]string, tdp *models.DelimiterPair, tlsSkipVerify bool) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]string, *models.DelimiterPair, bool) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMySQLQueryShowAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify
func (_m *mockAgentsRegistry) StartMySQLQueryShowAction(ctx context.Context, id string, pmmAgentID string, dsn string, query string, files map[string]string, tdp *models.DelimiterPair, tlsSkipVerify bool) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]string, *models.DelimiterPair, bool) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, query, files, tdp, tlsSkipVerify)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartPostgreSQLQuerySelectAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn, query
func (_m *mockAgentsRegistry) StartPostgreSQLQuerySelectAction(ctx context.Context, id string, pmmAgentID string, dsn string, query string) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartPostgreSQLQueryShowAction provides a mock function with given fields: ctx, id, pmmAgentID, dsn
func (_m *mockAgentsRegistry) StartPostgreSQLQueryShowAction(ctx context.Context, id string, pmmAgentID string, dsn string) error {
	ret := _m.Called(ctx, id, pmmAgentID, dsn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, id, pmmAgentID, dsn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockAgentsRegistry creates a new instance of mockAgentsRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAgentsRegistry(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAgentsRegistry {
	mock := &mockAgentsRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
