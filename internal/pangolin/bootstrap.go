package pangolin

import (
	"context"

	"github.com/TheCampground/tupai/internal/docker"
)

type PangolinClient struct {
	engine        *docker.DockerEngine
	containerName string
}

func New(engine *docker.DockerEngine, containerName string) *PangolinClient {
	return &PangolinClient{
		engine:        engine,
		containerName: containerName,
	}
}

// Bootstrap refers to first time setup of pangolin, this function automates it by automatically
// configuring the root user
func (client *PangolinClient) Bootstrap(ctx context.Context, email string, password string) error {
	return client.engine.ExecByName(ctx, client.containerName, []string{
		"pangctl",
		"set-admin-credentials",
		"--email",
		email,
		"--password",
		password,
	})
}
