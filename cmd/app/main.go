package main

import (
	"fmt"
	nats_adapter "websocket-gateway/infrastructure/nats"
	config_manager "websocket-gateway/pkg/config"

	"github.com/nats-io/nats.go"
)

func main() {
	configManager := config_manager.NewConfigManager()
	configManager.Bootstrap()

	natsUrl := configManager.Get("nats.url")
	natsUser := configManager.Get("nats.system_username")
	natsPassword := configManager.Get("nats.system_password")

	if (natsUrl != nil) && (natsUser != nil) && (natsPassword != nil) {
		natsConnection := nats_adapter.NewNatsConnection(natsUrl.(string), nats.Options{
			User:     natsUser.(string),
			Password: natsPassword.(string),
		})

		fmt.Println(natsConnection)
	}
}
