package main

import (
	"os"
	natsadapter "websocket-gateway/pkgs/brokers/nats"
	EnvManager "websocket-gateway/pkgs/env-manager"
	webserver "websocket-gateway/pkgs/web-server"

	"github.com/nats-io/nats.go"
)

func main() {
	EnvManager.Bootstrap()
	natsadapter.Connect(os.Getenv("NATS_URL"), nats.Options{User: os.Getenv("NATS_USERNAME"), Password: os.Getenv("NATS_PASSWORD")})
	webserver.Bootstrap()
}
