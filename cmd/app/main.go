package main

import (
	nats_adapter "websocket-gateway/infrastructure/brokers/nats"
	config_manager "websocket-gateway/pkg/config"
	"websocket-gateway/pkg/monitor"

	"github.com/nats-io/nats.go"
)

type Connection interface {
	Disconnect()
	Discover() any
}

func main() {
	configManager := config_manager.NewConfigManager()
	configManager.Bootstrap()

	broker := configManager.Get("broker")

	var connection Connection

	switch broker {
	case "nats":
		natsUrl := configManager.Get("nats.url")
		natsUser := configManager.Get("nats.system_username")
		natsPassword := configManager.Get("nats.system_password")

		if (natsUrl != nil) && (natsUser != nil) && (natsPassword != nil) {
			connection = nats_adapter.NewNatsConnection(natsUrl.(string), nats.Options{
				User:     natsUser.(string),
				Password: natsPassword.(string),
			})
		}
	}

	monitor := monitor.NewMonitor(connection.Discover)

	monitor.Discover()
}
