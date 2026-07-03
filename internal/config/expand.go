package config

import (
	"fmt"
	"os"
	"strings"
)

/*
Expands ENV template strings to real values

Currently is manual config traversal
*/
func (cfg *BootstrapConfig) Expand() error {
	if err := expandValue(&cfg.Version); err != nil {
		return err
	}

	if err := expandValue(&cfg.Container.Name); err != nil {
		return err
	}

	if err := expandValue(&cfg.RootAccount.Email); err != nil {
		return err
	}

	if err := expandValue(&cfg.RootAccount.Password); err != nil {
		return err
	}

	if err := expandValue(&cfg.Api.Url); err != nil {
		return err
	}

	return nil
}

func expandValue(template_str *string) error {
	if template_str == nil {
		return nil
	}

	if !isEnvTemplateStr(*template_str) {
		return nil
	}

	value, err := findEnvElement(stripTemplate(*template_str))

	if err != nil {
		return err
	}

	*template_str = *value
	return nil
}

func stripTemplate(key string) string {
	after, _ := strings.CutPrefix(key, "${")
	after, _ = strings.CutSuffix(after, "}")

	return after
}

func isEnvTemplateStr(key string) bool {
	if strings.HasPrefix(key, "${") && strings.HasSuffix(key, "}") {
		return true
	}

	return false
}

func findEnvElement(key string) (*string, error) {
	value, found := os.LookupEnv(key)

	if !found {
		return nil, fmt.Errorf("could not find env value with key: %s", key)
	}

	return &value, nil
}
