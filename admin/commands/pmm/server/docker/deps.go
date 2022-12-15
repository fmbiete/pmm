// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"context"
	"io"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/percona/pmm/admin/pkg/docker"
)

//go:generate ../../../../../bin/mockery -name=functions -case=snake -inpkg -testonly

// functions contain methods required to interact with Docker.
type functions interface {
	imager
	installer

	ChangeServerPassword(ctx context.Context, containerID, newPassword string) error
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
	ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
	CreateVolume(ctx context.Context, volumeName string, labels map[string]string) (*types.Volume, error)
	FindServerContainers(ctx context.Context) ([]types.Container, error)
	GetDockerClient() *client.Client
	RunContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (string, error)
	WaitForHealthyContainer(ctx context.Context, containerID string) <-chan docker.WaitHealthyResponse
}

// imager holds methods to interact with Docker images.
type imager interface {
	ParsePullImageProgress(r io.Reader, p *tea.Program) (<-chan struct{}, <-chan error)
	PullImage(ctx context.Context, dockerImage string, opts types.ImagePullOptions) (io.Reader, error)
}

// installer holds methods related to Docker installation.
type installer interface {
	HaveDockerAccess(ctx context.Context) bool
	InstallDocker(ctx context.Context) error
	IsDockerInstalled() (bool, error)
}
