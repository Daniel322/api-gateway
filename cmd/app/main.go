package main

import (
	config_manager "websocket-gateway/pkg/config"
)

func main() {
	configManager := config_manager.NewConfigManager()
	configManager.Bootstrap()
	// natsadapter.Connect(
	// 	os.Getenv("NATS_URL"),
	// 	nats.Options{
	// 		User:     os.Getenv("NATS_SYSTEM_USERNAME"),
	// 		Password: os.Getenv("NATS_SYSTEM_PASSWORD")})
	// webserver.Bootstrap()
}
