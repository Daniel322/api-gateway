package config_manager

import (
	"errors"
	"os"
	"slices"
)

func (manager *ConfigManager) isInvalidValue(value interface{}) bool {
	switch value.(type) {
	case string:
		return value == nil || len(value.(string)) == 0
	default:
		return value == nil
	}
}

// TODO: fix multi key for maps
// TODO: add tests

func (manager *ConfigManager) Get(key string) (interface{}, error) {
	if ok := slices.Contains(SUPPORTED_KEYS, key); !ok {
		return nil, errors.New("unsupported key " + key)
	}

	var value interface{}

	// go to map file
	value = manager.Values[key]
	manager.Logger.Println("get "+key+" from values"+" result:", value)
	if manager.isInvalidValue(value) {
		// go to os
		value = os.Getenv(key)
		manager.Logger.Println("get "+key+" from os"+" result:", value)
		if manager.isInvalidValue(value) {
			// try to write default value
			value = DEFAULT_CONFIG[key]
			manager.Logger.Println("get "+key+" from default"+" result:", value)
			if manager.isInvalidValue(value) {
				return nil, errors.New(key + " not found")
			}
		}
	}

	return value, nil
}
