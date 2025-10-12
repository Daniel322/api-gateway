package nats_adapter

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats.go"
)

type NatsConnection struct {
	Connection *nats.Conn
	logger     *log.Logger
}

func NewNatsConnection(url string, options nats.Options) *NatsConnection {
	logger := log.New(log.Writer(), "NatsConnection ", log.LstdFlags)
	conn, err := nats.Connect(url, nats.UserInfo(options.User, options.Password), nats.PermissionErrOnSubscribe(true))

	if err != nil {
		logger.Println("error on connect to NATS", err)
		return nil
	}

	logger.Println("succesfull connect to NATS!")

	return &NatsConnection{
		Connection: conn,
		logger:     logger,
	}
}

func Connect(url string, options nats.Options) error {
	_, err := nats.Connect(url, nats.UserInfo(options.User, options.Password), nats.PermissionErrOnSubscribe(true))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succesfull connect to NATS!")

	// natsConnection = nc
	// fmt.Println(methods)
	// _, err = natsConnection.Subscribe("$SRV.REGISTER", func(m *nats.Msg) {
	// 	fmt.Println(m.Subject)
	// 	var stringData string = string(m.Data)
	// 	fmt.Println(m.Subject, stringData)
	// 	var jsonData SrvRegisterMessage
	// 	json.Unmarshal(m.Data, &jsonData)
	// 	for _, v := range jsonData.Info.Endpoints {
	// 		fmt.Println("v", v)
	// 		if len(v.Metadata.Value) <= 0 && len(v.Metadata.ValueRegex) <= 0 {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("\n", jsonData)
	// })

	// if err != nil {
	// 	fmt.Println("subscribe on srv info error", err)
	// }
	return err
}

func (nats *NatsConnection) Disconnect() {
	nats.Connection.Close()
	nats.logger.Println("succesfull disconect from NATS")
}
