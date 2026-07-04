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

//go:embed testdata/organizations.yml
var organizationsFile []byte

func TestLoadConfigIncompatVer(t *testing.T) {
	_, err := loadFromBytes(invalidVerFile)

	if err == nil {
		t.Fatal("version should be invalid")
	}
}

func TestLoadConfigEnv(t *testing.T) {
	err := os.Setenv("TUPAI_ROOT_PASSWORD", "super-secret")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := loadFromBytes(envLoadFile)
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

func TestLoadConfigOrganizations(t *testing.T) {
	cfg, err := loadFromBytes(organizationsFile)
	if err != nil {
		t.Fatal(err)
	}

	err = cfg.Expand()
	if err != nil {
		t.Fatal(err)
	}

	org := cfg.Organizations[0]

	if org.ID != "thecampground" {
		t.Fatalf("org id was incorrect")
	}

	if org.Name != "The Campground" {
		t.Fatalf("org name was incorrect")
	}

	if org.Subnet != "100.90.128.0/20" {
		t.Fatalf("org subnet was incorrect")
	}

	if org.Utility != "100.96.128.0/20" {
		t.Fatalf("org utility subnet was incorrect")
	}
}

func TestLoadFromBytesInvalidYAML(t *testing.T) {
	_, err := loadFromBytes([]byte("not: valid: yaml: at: all: :("))
	if err == nil {
		t.Fatal("expected error for malformed yaml")
	}
}
