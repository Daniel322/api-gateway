package natsadapter

import (
	"encoding/json"
	"fmt"
	"reflect"

	nats "github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn

var methods map[string]any

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url)
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc

	natsConnection.Subscribe("$SRV.REGISTER", func(m *nats.Msg) {
		// var stringData string = string(m.Data)
		var jsonData map[string]map[string]any
		json.Unmarshal(m.Data, &jsonData)
		for key, value := range jsonData {
			if reflect.TypeOf(value).Kind().String() == "map" {
				for infoKey, infoValue := range value {
					fmt.Println("key:", infoKey, "value:", infoValue)
				}
			} else {
				fmt.Println("key:", key, "value:", value)
			}
		}
		fmt.Println(jsonData)
	})

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
