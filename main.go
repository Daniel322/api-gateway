package main

import (
	EnvManager "websocket-gateway/pkgs/env-manager"
	webserver "websocket-gateway/pkgs/web-server"
)

func main() {
	EnvManager.Bootstrap()
	webserver.Bootstrap()
}
