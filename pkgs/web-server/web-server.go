package webserver

import (
	"net/http"
	"os"
	wsconnection "websocket-gateway/pkgs/ws-connection"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Bootstrap() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to WS Gateway!")
	})
	e.GET("/ws", wsconnection.CreateConnection)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
