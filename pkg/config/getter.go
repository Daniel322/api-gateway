package config_manager

import (
	"errors"
	"os"
	"slices"
	"strings"
)

func (manager *ConfigManager) isInvalidValue(value interface{}) bool {
	switch value.(type) {
	case string:
		return value == nil || len(value.(string)) == 0
	default:
		return value == nil
	}
}

func baseGet(key string, value map[string]interface{}) interface{} {
	keyParts := strings.Split(key, ".")
	result := value
	for _, part := range keyParts {
		if stepValue, ok := result[part]; ok {
			switch stepValueType := stepValue.(type) {
			case map[string]interface{}:
				result = stepValueType
			default:
				return stepValueType
			}
		} else {
			return nil
		}
	}
	return result
}

// TODO: add tests

func (manager *ConfigManager) Get(key string) (interface{}, error) {
	if ok := slices.Contains(SUPPORTED_KEYS, key); !ok {
		return nil, errors.New("unsupported key " + key)
	}

	var value interface{}

	value = baseGet(key, manager.Values)
	manager.Logger.Println("get "+key+" from values"+" result:", value)

	if manager.isInvalidValue(value) {
		// go to os
		value = os.Getenv(key)
		manager.Logger.Println("get "+key+" from os"+" result:", value)
		if manager.isInvalidValue(value) {
			// try to write default value
			value = baseGet(key, DEFAULT_CONFIG)
			manager.Logger.Println("get "+key+" from default"+" result:", value)

			if manager.isInvalidValue(value) {
				return nil, errors.New(key + " not found")
			}
		}
	}

	return value, nil
}
