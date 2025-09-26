package config_manager

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigWebServer struct {
	Port int
}

type ConfigManager struct {
	Name   string
	Logger *log.Logger
	// TODO: make possibility to keep vars here, read file and set vars in custom map
	Values map[string]string
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Name:   "ConfigManager",
		Logger: log.New(os.Stdout, "ConfigManager ", log.LstdFlags),
		Values: make(map[string]string),
	}
}

func (manager *ConfigManager) Bootstrap() {
	manager.Logger.Println("start config bootstrap")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
