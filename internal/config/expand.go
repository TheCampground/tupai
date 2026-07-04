package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Expands TupaiConfig in-place
func (cfg *TupaiConfig) Expand() error {
	return walkConfigRecursively(reflect.ValueOf(cfg).Elem())
}

func walkConfigRecursively(v reflect.Value) error {
	switch v.Kind() {
	case reflect.Pointer:
		return walkPointer(v)
	case reflect.Struct:
		return walkStruct(v)
	case reflect.Map:
		return walkMap(v)
	case reflect.Slice, reflect.Array:
		return walkSliceOrArray(v)
	case reflect.String:
		return expandStringValue(v)
	}

	return nil
}

func walkPointer(ptr reflect.Value) error {
	if !ptr.IsNil() {
		return walkConfigRecursively(ptr.Elem())
	}

	return nil
}

func walkStruct(strct reflect.Value) error {
	for _, field := range strct.Fields() {
		if err := walkConfigRecursively(field); err != nil {
			return err
		}
	}

	return nil
}

func walkMap(mp reflect.Value) error {
	for _, key := range mp.MapKeys() {
		val := mp.MapIndex(key)

		newVal := reflect.New(val.Type()).Elem()
		newVal.Set(val)

		if err := walkConfigRecursively(newVal); err != nil {
			return err
		}

		// Potentially dangerous approach, will not work with non-str map keys
		mp.SetMapIndex(key, newVal)
	}

	return nil
}

func walkSliceOrArray(list reflect.Value) error {
	for i := 0; i < list.Len(); i++ {
		if err := walkConfigRecursively(list.Index(i)); err != nil {
			return err
		}
	}

	return nil
}

func expandStringValue(str reflect.Value) error {
	if !str.CanSet() {
		return nil
	}

	expandedStr, err := expandString(str.String())

	if err != nil {
		return err
	}

	if expandedStr == nil {
		return nil
	}

	str.SetString(*expandedStr)

	return nil
}

func expandString(str string) (*string, error) {
	if !isEnvTemplateStr(str) {
		return nil, nil
	}

	value, err := findEnvElement(stripTemplate(str))

	if err != nil {
		return nil, err
	}

	return value, nil
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
