package main

import (
	config_manager "websocket-gateway/pkg/config"
)

func main() {
	configManager := config_manager.NewConfigManager()
	configManager.Bootstrap()
}
