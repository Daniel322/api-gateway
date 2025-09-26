package envmanager

import (
	"fmt"

	"github.com/joho/godotenv"
)

type ConfigWebServer struct {
	Port int
}

type EnvManager struct {
	Name string
	// TODO: make possibility to keep vars here, read file and set vars in custom map
	Values map[string]string
}

func NewEnvManager() *EnvManager {
	return &EnvManager{
		Name:   "env_manager",
		Values: make(map[string]string),
	}
}

func (manager *EnvManager) Bootstrap() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
