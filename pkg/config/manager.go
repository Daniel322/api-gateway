package config_manager

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type ConfigWebServer struct {
	Port int
}

type Config struct {
	Server ConfigWebServer
}

type ConfigManager struct {
	Name   string
	Logger *log.Logger
	// TODO: make possibility to keep vars here, read file and set vars in custom map
	Values Config
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Name:   "ConfigManager",
		Logger: log.New(os.Stdout, "ConfigManager ", log.LstdFlags),
		Values: Config{},
	}
}

func (manager *ConfigManager) Bootstrap(path string) {
	manager.Logger.Println("start config bootstrap")

	absPath, err := filepath.Abs(path)

	if err != nil {
		manager.Logger.Println("error on get absolute path", err)
	}

	manager.Logger.Println("absolute path", absPath)

	bytes, err := os.ReadFile(absPath)

	if err != nil {
		manager.Logger.Println("error on read file", err)
	}

	err = json.Unmarshal(bytes, &manager.Values)

	if err != nil {
		manager.Logger.Println("error on parse bytes", err)
	}
}
