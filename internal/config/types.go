package config

var (
	VERSIONS = []string{"0.0.1"}
)

type BootstrapConfig struct {
	Version     string      `yaml:"version"`
	Container   Container   `yaml:"container"`
	RootAccount RootAccount `yaml:"rootAccount"`
	Api         Api         `yaml:"api"`
}

type Container struct {
	Name string `yaml:"name"`
}

type RootAccount struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

type Api struct {
	Url          string `yaml:"url"`
	CreateApiKey bool   `yaml:"createApiKey"`
}
