package config

import (
	"fmt"
	"os"
	"slices"

	"github.com/goccy/go-yaml"
)

func LoadBoostrap(path string) (*BootstrapConfig, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return loadFromBytes(&file)
}

func validateVersion(cfg *BootstrapConfig) error {
	if slices.Contains(VERSIONS, cfg.Version) {
		return nil
	}

	return fmt.Errorf("current version: %s is not supported, supported versions: %v", cfg.Version, VERSIONS)
}

func loadFromBytes(str *[]byte) (*BootstrapConfig, error) {
	var cfg BootstrapConfig

	if err := yaml.Unmarshal(*str, &cfg); err != nil {
		return nil, err
	}

	if err := validateVersion(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
