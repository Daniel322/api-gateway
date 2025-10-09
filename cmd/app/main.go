package main

import (
	"fmt"
	config_manager "websocket-gateway/pkg/config"
)

func main() {
	configManager := config_manager.NewConfigManager()
	configManager.Bootstrap()

	fmt.Printf("%+v\n", configManager.Values)
	test1, err := configManager.Get("server.port")
	if err != nil {
		fmt.Println("fail test1", err)
	} else {
		fmt.Println("complete test1", test1)
	}

	test3, err := configManager.Get("testInclude.asd")
	if err != nil {
		fmt.Println("fail test3", err)
	} else {
		fmt.Println("complete test3", test3)
	}
	// fmt.Printf("%+v\n", os.Environ())

	// natsadapter.Connect(
	// 	os.Getenv("NATS_URL"),
	// 	nats.Options{
	// 		User:     os.Getenv("NATS_SYSTEM_USERNAME"),
	// 		Password: os.Getenv("NATS_SYSTEM_PASSWORD")})
	// webserver.Bootstrap()
}
