package natsadapter

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn

var methods map[string]any

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url, nats.UserInfo(options.User, options.Password), nats.PermissionErrOnSubscribe(true))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc

	_, err = natsConnection.Subscribe("$SYS.>", func(m *nats.Msg) {
		fmt.Println(m.Subject)
		var stringData string = string(m.Data)
		fmt.Println(m.Subject, "stringData", stringData)
		var jsonData map[string]map[string]any
		json.Unmarshal(m.Data, &jsonData)
		fmt.Println(m.Subject, "jsonData", jsonData)
	})

	if err != nil {
		fmt.Println(err)
	}

	_, err = natsConnection.Subscribe("$SRV.>", func(m *nats.Msg) {
		fmt.Println(m.Subject)
		var stringData string = string(m.Data)
		fmt.Println(m.Subject, stringData)
		var jsonData map[string]map[string]any
		json.Unmarshal(m.Data, &jsonData)
		fmt.Println(m.Subject, "jsonData", jsonData)
	})

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
