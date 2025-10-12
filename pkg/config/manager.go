package config_manager

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: add yaml support
// TODO: add .env support

type ConfigManager struct {
	logger *log.Logger
	values map[string]interface{}
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		logger: log.New(os.Stdout, "ConfigManager ", log.LstdFlags),
		values: make(map[string]interface{}),
	}
}

type CheckFileFormatResult struct {
	Filename string
	Format   string
}

func (manager *ConfigManager) checkFileFormat() (*CheckFileFormatResult, error) {
	absPath, err := filepath.Abs(CONFIG_FOLDER)

	if err != nil {
		manager.logger.Println("error on get absolute path", err)
		return nil, err
	}

	files, err := os.ReadDir(absPath)

	if err != nil {
		manager.logger.Println("error on read abs dir", err)
		return nil, err
	}

	configFile := files[0].Name()

	return &CheckFileFormatResult{
		Filename: CONFIG_FOLDER + "/" + configFile,
		Format:   strings.Split(configFile, ".")[1],
	}, nil
}

func (manager *ConfigManager) Bootstrap() {
	manager.logger.Println("start config bootstrap")

	checkFormatResult, err := manager.checkFileFormat()

	manager.logger.Println(checkFormatResult)

	if err != nil {
		manager.logger.Println("error on get fileformat", err)
	}

	absPath, err := filepath.Abs(checkFormatResult.Filename)

	if err != nil {
		manager.logger.Println("error on get absolute path", err)
	}

	manager.logger.Println("absolute path", absPath)
	manager.logger.Println("file format", checkFormatResult.Format)

	bytes, err := os.ReadFile(absPath)

	if err != nil {
		manager.logger.Println("error on read file", err)
	}

	switch checkFormatResult.Format {
	case "json":
		manager.setJSON(bytes)
	case "env":
		manager.setEnv(bytes)
	case "yaml":
		manager.setYaml(bytes)
	default:
		manager.logger.Panicln("unsupported config file type")
	}

	manager.setToOS()
}

func (manager *ConfigManager) setEnv(bytes []byte) {
	manager.logger.Panicln("env config not implemented")
}

func (manager *ConfigManager) setYaml(bytes []byte) {
	manager.logger.Panicln("yaml config not implemented")
}

func (manager *ConfigManager) setJSON(bytes []byte) {
	err := json.Unmarshal(bytes, &manager.values)

	if err != nil {
		manager.logger.Println("error on parse bytes in map", err)
	}
}

func (manager *ConfigManager) setToOS() {
	manager.logger.Println("Start to backup write config to OS")
	for key, value := range manager.values {
		switch v := value.(type) {
		case map[string]interface{}:
			for k, val := range v {
				manager.baseSetToOS(key+"."+k, val)
			}
		default:
			manager.baseSetToOS(key, value)
		}
	}
	manager.logger.Println("Complete to backup write config to OS")
}

func (manager *ConfigManager) baseSetToOS(key string, value interface{}) {
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
