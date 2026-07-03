package docker

import (
	"context"
	"testing"
	"time"

	"github.com/docker/go-sdk/client"
	"github.com/docker/go-sdk/container"
)

func TestDockerClientCreate(t *testing.T) {
	cli, err := NewClient(t.Context())

	if err != nil {
		t.Fatalf("could not create docker client: %s", err)
	}

	if err = cli.Close(); err != nil {
		t.Fatalf("could not close client: %s", err)
	}
}

func TestDockerClientExecByContainer(t *testing.T) {
	cli, err := NewClient(t.Context())
	if err != nil {
		t.Fatalf("could not create docker client: %s", err)
	}
	defer cli.Close()

	ctr, err := startAlpineContainer(t.Context(), cli, "test")
	if err != nil {
		t.Fatalf("Could not start test container: %s", err)
	}
	defer ctr.Terminate(t.Context(), container.TerminateTimeout(time.Millisecond*1))

	engine := NewEngine(cli)
	if err = engine.execByContainer(t.Context(), ctr, []string{"echo", "hello"}); err != nil {
		t.Fatalf("could not exec: %s", err)
	}
}

func TestDockerClientExecByName(t *testing.T) {
	cli, err := NewClient(t.Context())
	if err != nil {
		t.Fatalf("could not create docker client: %s", err)
	}
	defer cli.Close()

	ctr, err := startAlpineContainer(t.Context(), cli, "test")
	if err != nil {
		t.Fatalf("Could not start test container: %s", err)
	}
	defer ctr.Terminate(t.Context(), container.TerminateTimeout(time.Millisecond*1))

	engine := NewEngine(cli)
	if err = engine.ExecByName(t.Context(), "test", []string{"echo", "hello"}); err != nil {
		t.Fatalf("could not exec: %s", err)
	}
}

func startAlpineContainer(ctx context.Context, cli client.SDKClient, name string) (*container.Container, error) {
	contain, err := container.Run(
		ctx,
		container.WithImage("alpine:latest"),
		container.WithClient(cli),
		container.WithCredentialsFn(func(registry string) (string, string, error) {
			return "", "", nil
		}),
		container.WithCmd("sleep", "infinity"),
		container.WithName(name),
	)
	if err != nil {
		return nil, err
	}

	return contain, nil
}
