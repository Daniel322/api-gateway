package config_manager

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
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

const CONFIG_FOLDER = "./config"

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Name:   "ConfigManager",
		Logger: log.New(os.Stdout, "ConfigManager ", log.LstdFlags),
		Values: Config{},
	}
}

type CheckFileFormatResult struct {
	Filename string
	Format   string
}

func (manager *ConfigManager) CheckFileFormat() (*CheckFileFormatResult, error) {
	absPath, err := filepath.Abs(CONFIG_FOLDER)

	if err != nil {
		manager.Logger.Println("error on get absolute path", err)
		return nil, err
	}

	files, err := os.ReadDir(absPath)

	if err != nil {
		manager.Logger.Println("error on read abs dir", err)
		return nil, err
	}

	configFile := files[0].Name()

	manager.Logger.Println("configFile", configFile)

	return &CheckFileFormatResult{
		Filename: CONFIG_FOLDER + "/" + configFile,
		Format:   strings.Split(configFile, ".")[1],
	}, nil
}

func (manager *ConfigManager) Bootstrap() {
	manager.Logger.Println("start config bootstrap")

	checkFormatResult, err := manager.CheckFileFormat()

	if err != nil {
		manager.Logger.Println("error on get fileformat", err)
	}

	absPath, err := filepath.Abs(checkFormatResult.Filename)

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
