package natsadapter

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn

var methods map[string]any

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url, nats.Name(options.User))
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc

	nc.Subscribe("$SYS.> ", func(m *nats.Msg) {
		fmt.Println(m.Subject)
		var stringData string = string(m.Data)
		fmt.Println(stringData)
		var jsonData map[string]map[string]any
		json.Unmarshal(m.Data, &jsonData)
		fmt.Println(jsonData)
	})

	// natsConnection.Subscribe("$SYS.SERVER.*.STATSZ ", func(m *nats.Msg) {
	// 	fmt.Println("STATS")
	// 	var stringData string = string(m.Data)
	// 	fmt.Println(stringData)
	// 	var jsonData map[string]map[string]any
	// 	json.Unmarshal(m.Data, &jsonData)
	// 	fmt.Println(jsonData)
	// })

	// natsConnection.Subscribe("$SYS.REQ.SERVER.PING", func(m *nats.Msg) {
	// 	fmt.Println("PING")
	// 	var stringData string = string(m.Data)
	// 	fmt.Println(stringData)
	// 	var jsonData map[string]map[string]any
	// 	json.Unmarshal(m.Data, &jsonData)
	// 	fmt.Println(jsonData)
	// })

	// natsConnection.Subscribe("$SYS.SERVER.*.CLIENT.CONNECT", func(m *nats.Msg) {
	// 	fmt.Println("connect \n")
	// 	var stringData string = string(m.Data)
	// 	fmt.Println(stringData, "\n")
	// 	var jsonData map[string]map[string]any
	// 	json.Unmarshal(m.Data, &jsonData)
	// 	fmt.Println(jsonData)
	// 	// for key, value := range jsonData {
	// 	// 	if reflect.TypeOf(value).Kind().String() == "map" {
	// 	// 		for infoKey, infoValue := range value {
	// 	// 			fmt.Println("key:", infoKey, "value:", infoValue)
	// 	// 		}
	// 	// 	} else {
	// 	// 		fmt.Println("key:", key, "value:", value)
	// 	// 	}
	// 	// }
	// })

	// natsConnection.Subscribe("$SYS.SERVER.*.CLIENT.DISCONNECT", func(message *nats.Msg) {
	// 	fmt.Println("disconnect \n")
	// 	var stringData string = string(message.Data)
	// 	fmt.Println(stringData, "\n")
	// 	var jsonData map[string]map[string]any
	// 	json.Unmarshal(message.Data, &jsonData)
	// 	fmt.Println(jsonData)
	// })

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
