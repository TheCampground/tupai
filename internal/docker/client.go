package docker

import (
	"context"

	"github.com/docker/go-sdk/client"
)

func NewClient(ctx context.Context) (client.SDKClient, error) {
	cli, err := client.New(
		ctx,
		client.WithDockerHost("unix:///var/run/docker.sock"),
	)

	if err != nil {
		return nil, err
	}

	return cli, nil
}
