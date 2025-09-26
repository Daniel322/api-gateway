package wsconnection

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"websocket-gateway/pkg/utils"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024, // Размер буфера чтения
		WriteBufferSize: 1024, // Размер буфера записи
		// Позволяет определить, должен ли сервер сжимать сообщения
		EnableCompression: true,
	}
)

func CreateConnection(c echo.Context) error {
	var connectionId = utils.GenerateNewId()

	keepalive, err := strconv.Atoi(os.Getenv("KEEPALIVE_TIME"))
	if err != nil {
		c.Logger().Error(err)
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	ws.WriteMessage(websocket.TextMessage, []byte("your connection id"+" "+connectionId))
	stop := utils.SetInterval(func() {
		fmt.Println(connectionId + " interval cb")
		ws.WriteMessage(websocket.TextMessage, []byte(connectionId+" keepalive"))
	}, time.Duration(keepalive)*time.Second)
	if err != nil {
		stop <- true
		c.Logger().Error(err)
	}
	closeHandler := ws.CloseHandler()
	ws.SetCloseHandler(func(code int, text string) error {
		stop <- true
		fmt.Printf("Connection "+connectionId+" closed with code %d and text: %s\n", code, text)
		err = closeHandler(code, text)
		return err
	})
	defer ws.Close()

	for {
		// Read
		var msgData WsMessage[CallMessage]
		_, msg, err := ws.ReadMessage()
		json.Unmarshal(msg, &msgData)
		if err != nil {
			stop <- true
			c.Logger().Error(err)
			break
		}
		fmt.Printf("%s\n", msgData.Message)
		err = ws.WriteMessage(websocket.TextMessage, []byte(connectionId+" receive "+string(msgData.Message.CallId)))
		if err != nil {
			stop <- true
			c.Logger().Error(err)
			break
		}
	}

	return err
}
