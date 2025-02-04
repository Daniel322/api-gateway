package EnvManager

import (
	"fmt"

	"github.com/joho/godotenv"
)

var err error

func Bootstrap() {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
