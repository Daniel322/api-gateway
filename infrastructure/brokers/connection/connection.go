package connection

import "log"

type Connection[conn any] struct {
	Connection conn
	Logger     *log.Logger
}

func (Connection[conn]) Disconnect() {}
