package bootstrap

import (
	"context"
	"fmt"

	"github.com/TheCampground/tupai/internal/config"
	"github.com/TheCampground/tupai/internal/docker"
	"github.com/TheCampground/tupai/internal/pangolin"
)

func LoadAndBoostrap(ctx context.Context, configPath string) error {
	cfg, err := config.LoadBoostrap(configPath)
	if err != nil {
		return fmt.Errorf("failed to read boostrap config: %w", err)
	}

	if err = cfg.Expand(); err != nil {
		return fmt.Errorf("failed to process environment template strings: %w", err)
	}

	dockerClient, err := docker.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to docker daemon socket: %w", err)
	}
	defer dockerClient.Close()

	dockerEngine := docker.NewEngine(dockerClient)
	pangolinClient := pangolin.New(dockerEngine, cfg.Container.Name)

	err = pangolinClient.Bootstrap(ctx, cfg.RootAccount.Email, cfg.RootAccount.Password)
	if err != nil {
		return fmt.Errorf("failed to set admin credentials: %w", err)
	}

	return nil
}
