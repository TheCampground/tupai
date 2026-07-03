package config

import (
	_ "embed"
	"os"
	"testing"
)

//go:embed testdata/env_load.yml
var envLoadFile []byte

//go:embed testdata/invalid_ver.yml
var invalidVerFile []byte

func TestLoadConfigIncompatVer(t *testing.T) {
	_, err := loadFromBytes(&invalidVerFile)

	if err == nil {
		t.Fatal("version should be invalid")
	}
}

func TestLoadConfigEnv(t *testing.T) {
	err := os.Setenv("TUPAI_ROOT_PASSWORD", "super-secret")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := loadFromBytes(&envLoadFile)
	if err != nil {
		t.Fatal(err)
	}

	err = cfg.Expand()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.RootAccount.Password != "super-secret" {
		t.Fatal("root account password does not match")
	}
}
