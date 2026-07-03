package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/go-sdk/client"
	"github.com/docker/go-sdk/container"
	mContainer "github.com/moby/moby/api/types/container"
)

type DockerEngine struct {
	cli client.SDKClient
}

func NewEngine(cli client.SDKClient) *DockerEngine {
	return &DockerEngine{
		cli: cli,
	}
}

func (e *DockerEngine) FindByName(ctx context.Context, containerName string) (*mContainer.Summary, *container.Container, error) {
	summary, err := e.cli.FindContainerByName(ctx, containerName)

	if err != nil {
		return nil, nil, fmt.Errorf("container %s not found: %w", containerName, err)
	}

	ctr, err := container.FromID(ctx, e.cli, summary.ID)
	if err != nil {
		return nil, nil, err
	}

	return summary, ctr, nil
}

func (e *DockerEngine) ExecByName(ctx context.Context, containerName string, cmd []string) error {
	summary, ctr, err := e.FindByName(ctx, containerName)

	if err != nil {
		return fmt.Errorf("container %s not found: %w", containerName, err)
	}

	if summary.Health.Status != "healthy" && summary.Health.Status != "none" {
		return fmt.Errorf("container status is: %s expected healthly", summary.Health.Status)
	}

	return e.execByContainer(ctx, ctr, cmd)
}

func (e *DockerEngine) execByContainer(ctx context.Context, ctr *container.Container, cmd []string) error {
	status, read, err := ctr.Exec(ctx, cmd)

	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	go func() {
		_, _ = io.Copy(os.Stdout, read)
	}()

	if status != 0 {
		return fmt.Errorf("command failed with exit-code: %d", status)
	}

	return nil
}
