package wsconnection

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsConnection struct {
	ConnectionId     string                     `json:"connectionId"`
	CreateConnection func(c echo.Context) error `json:"createConnection"`
	Websocket        *websocket.Conn            `json:"websocket"`
}

type WsMessageType int64

const (
	Message WsMessageType = iota
	Subscribe
	Unsubscribe
)

func (t WsMessageType) String() string {
	switch t {
	case 0:
		return "message"
	case 1:
		return "subscribe"
	case 2:
		return "unsubscribe"
	}
	return "unsupported message type"
}

type CallMessage struct {
	Method string      `json:"method"`
	CallId string      `json:"callId"`
	Params interface{} `json:"params"`
}

type ResultMessage struct {
	CallId string      `json:"callId"`
	Result interface{} `json:"result"`
}

type ErrorMessage struct {
	Code   int64  `json:"code"`
	CallId string `json:"callId"`
	Error  string `json:"error"`
}

type WsMessage[T comparable] struct {
	Type    WsMessageType `json:"type"`
	Message T             `json:"message"`
}
