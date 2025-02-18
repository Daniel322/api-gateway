package natsadapter

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn
var err error

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url)
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc

	natsConnection.Subscribe("$SRV.REGISTER", func(m *nats.Msg) {
		var stringData string = string(m.Data)
		var jsonData any
		json.Unmarshal(m.Data, &jsonData)
		fmt.Println(jsonData)
		fmt.Println(stringData)
	})

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
