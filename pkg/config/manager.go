package config_manager

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: add possibility to merge with default values
// TODO: add yaml support
// TODO: add .env support

type ConfigWebServer struct {
	Port int
}

type Config struct {
	Server ConfigWebServer
}

type ConfigManager struct {
	Name   string
	Logger *log.Logger
	Values map[string]interface{}
}

const CONFIG_FOLDER = "./config"

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Name:   "ConfigManager",
		Logger: log.New(os.Stdout, "ConfigManager ", log.LstdFlags),
		Values: make(map[string]interface{}),
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

	return &CheckFileFormatResult{
		Filename: CONFIG_FOLDER + "/" + configFile,
		Format:   strings.Split(configFile, ".")[1],
	}, nil
}

func (manager *ConfigManager) Bootstrap() {
	manager.Logger.Println("start config bootstrap")

	checkFormatResult, err := manager.CheckFileFormat()

	manager.Logger.Println(checkFormatResult)

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

	switch checkFormatResult.Format {
	case "json":
		{
			manager.SetJSON(bytes)
		}
	}

	manager.SetToOS()
}

func (manager *ConfigManager) SetJSON(bytes []byte) {
	err := json.Unmarshal(bytes, &manager.Values)

	if err != nil {
		manager.Logger.Println("error on parse bytes in map", err)
	}
}

func (manager *ConfigManager) SetToOS() {
	manager.Logger.Println("Start to backup write config to OS")
	for key, value := range manager.Values {
		switch v := value.(type) {
		case map[string]interface{}:
			for k, val := range v {
				manager.baseSet(key+"."+k, val)
			}
		default:
			manager.baseSet(key, value)
		}
	}
	manager.Logger.Println("Complete to backup write config to OS")
}

func (manager *ConfigManager) baseSet(key string, value interface{}) {
	switch v := value.(type) {
	case string:
		os.Setenv(key, v)
	case float64:
		os.Setenv(key, strconv.FormatFloat(value.(float64), 'g', -1, 64))
	case bool:
		if v {
			os.Setenv(key, "true")
		} else {
			os.Setenv(key, "false")
		}
	}
}
