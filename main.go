package main

import (
	"net/http"
	EnvManager "websocket-gateway/pkgs/env-manager"
	wsconnection "websocket-gateway/pkgs/ws-connection"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	EnvManager.Bootstrap()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", wsconnection.CreateConnection)
	e.Logger.Fatal(e.Start(":1323"))
}
