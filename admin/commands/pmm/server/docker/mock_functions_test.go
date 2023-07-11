// Code generated by mockery v2.31.1. DO NOT EDIT.

package docker

import (
	context "context"
	io "io"

	tea "github.com/charmbracelet/bubbletea"
	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	volume "github.com/docker/docker/api/types/volume"
	client "github.com/docker/docker/client"
	mock "github.com/stretchr/testify/mock"

	pkgdocker "github.com/percona/pmm/admin/pkg/docker"
)

// MockFunctions is an autogenerated mock type for the Functions type
type MockFunctions struct {
	mock.Mock
}

// ChangeServerPassword provides a mock function with given fields: ctx, containerID, newPassword
func (_m *MockFunctions) ChangeServerPassword(ctx context.Context, containerID string, newPassword string) error {
	ret := _m.Called(ctx, containerID, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, containerID, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerInspect provides a mock function with given fields: ctx, containerID
func (_m *MockFunctions) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, containerID)

	var r0 types.ContainerJSON
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (types.ContainerJSON, error)); ok {
		return rf(ctx, containerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, containerID)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerStop provides a mock function with given fields: ctx, containerID, timeout
func (_m *MockFunctions) ContainerStop(ctx context.Context, containerID string, timeout *int) error {
	ret := _m.Called(ctx, containerID, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *int) error); ok {
		r0 = rf(ctx, containerID, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerUpdate provides a mock function with given fields: ctx, containerID, updateConfig
func (_m *MockFunctions) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	ret := _m.Called(ctx, containerID, updateConfig)

	var r0 container.ContainerUpdateOKBody
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, container.UpdateConfig) (container.ContainerUpdateOKBody, error)); ok {
		return rf(ctx, containerID, updateConfig)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, container.UpdateConfig) container.ContainerUpdateOKBody); ok {
		r0 = rf(ctx, containerID, updateConfig)
	} else {
		r0 = ret.Get(0).(container.ContainerUpdateOKBody)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, container.UpdateConfig) error); ok {
		r1 = rf(ctx, containerID, updateConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerWait provides a mock function with given fields: ctx, containerID, condition
func (_m *MockFunctions) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ret := _m.Called(ctx, containerID, condition)

	var r0 <-chan container.WaitResponse
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context, string, container.WaitCondition) (<-chan container.WaitResponse, <-chan error)); ok {
		return rf(ctx, containerID, condition)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, container.WaitCondition) <-chan container.WaitResponse); ok {
		r0 = rf(ctx, containerID, condition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan container.WaitResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, container.WaitCondition) <-chan error); ok {
		r1 = rf(ctx, containerID, condition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: ctx, volumeName, labels
func (_m *MockFunctions) CreateVolume(ctx context.Context, volumeName string, labels map[string]string) (*volume.Volume, error) {
	ret := _m.Called(ctx, volumeName, labels)

	var r0 *volume.Volume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) (*volume.Volume, error)); ok {
		return rf(ctx, volumeName, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *volume.Volume); ok {
		r0 = rf(ctx, volumeName, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.Volume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, volumeName, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindServerContainers provides a mock function with given fields: ctx
func (_m *MockFunctions) FindServerContainers(ctx context.Context) ([]types.Container, error) {
	ret := _m.Called(ctx)

	var r0 []types.Container
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]types.Container, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []types.Container); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDockerClient provides a mock function with given fields:
func (_m *MockFunctions) GetDockerClient() *client.Client {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
		}
	}

	return r0
}

// HaveDockerAccess provides a mock function with given fields: ctx
func (_m *MockFunctions) HaveDockerAccess(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InstallDocker provides a mock function with given fields: ctx
func (_m *MockFunctions) InstallDocker(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsDockerInstalled provides a mock function with given fields:
func (_m *MockFunctions) IsDockerInstalled() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParsePullImageProgress provides a mock function with given fields: r, p
func (_m *MockFunctions) ParsePullImageProgress(r io.Reader, p *tea.Program) (<-chan struct{}, <-chan error) {
	ret := _m.Called(r, p)

	var r0 <-chan struct{}
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(io.Reader, *tea.Program) (<-chan struct{}, <-chan error)); ok {
		return rf(r, p)
	}
	if rf, ok := ret.Get(0).(func(io.Reader, *tea.Program) <-chan struct{}); ok {
		r0 = rf(r, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(io.Reader, *tea.Program) <-chan error); ok {
		r1 = rf(r, p)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// PullImage provides a mock function with given fields: ctx, dockerImage, opts
func (_m *MockFunctions) PullImage(ctx context.Context, dockerImage string, opts types.ImagePullOptions) (io.Reader, error) {
	ret := _m.Called(ctx, dockerImage, opts)

	var r0 io.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) (io.Reader, error)); ok {
		return rf(ctx, dockerImage, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) io.Reader); ok {
		r0 = rf(ctx, dockerImage, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r1 = rf(ctx, dockerImage, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunContainer provides a mock function with given fields: ctx, config, hostConfig, containerName
func (_m *MockFunctions) RunContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (string, error) {
	ret := _m.Called(ctx, config, hostConfig, containerName)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, string) (string, error)); ok {
		return rf(ctx, config, hostConfig, containerName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, string) string); ok {
		r0 = rf(ctx, config, hostConfig, containerName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, string) error); ok {
		r1 = rf(ctx, config, hostConfig, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForHealthyContainer provides a mock function with given fields: ctx, containerID
func (_m *MockFunctions) WaitForHealthyContainer(ctx context.Context, containerID string) <-chan pkgdocker.WaitHealthyResponse {
	ret := _m.Called(ctx, containerID)

	var r0 <-chan pkgdocker.WaitHealthyResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan pkgdocker.WaitHealthyResponse); ok {
		r0 = rf(ctx, containerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan pkgdocker.WaitHealthyResponse)
		}
	}

	return r0
}

// NewMockFunctions creates a new instance of MockFunctions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFunctions(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockFunctions {
	mock := &MockFunctions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
