package natsadapter

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn
var err error

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url)
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc

	return err
}

func Disconnect() error {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
