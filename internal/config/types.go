package config

var (
	VERSIONS = []string{"0.0.1"}
)

type TupaiConfig struct {
	Version       string         `yaml:"version"`
	Container     Container      `yaml:"container"`
	RootAccount   RootAccount    `yaml:"rootAccount"`
	Organizations []Organization `yaml:"organizations"`
	Api           Api            `yaml:"api"`
}

type Container struct {
	Name string `yaml:"name"`
}

type RootAccount struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

type Organization struct {
	ID      string `yaml:"id"`
	Name    string `yaml:"name"`
	Subnet  string `yaml:"subnet"`
	Utility string `yaml:"utility"`
}

type Api struct {
	Url          string `yaml:"url"`
	CreateApiKey bool   `yaml:"createApiKey"`
}
