package natsadapter

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type SrvRegisterInfoEndpointMetadata struct {
	Domain     string `json:"domain"`
	Value      string `json:"value"`
	ValueRegex string `json:"valueRegex"`
	Field      string `json:"field"`
	AuthRole   string `json:"authRole"`
}

type SrvRegisterInfoEndpoints struct {
	Name     string                          `json:"name"`
	Subject  string                          `json:"subject"`
	Metadata SrvRegisterInfoEndpointMetadata `json:"metadata"`
}

type SrvRegisterInfo struct {
	Name      string                     `json:"name"`
	Id        string                     `json:"id"`
	Version   string                     `json:"version"`
	Endpoints []SrvRegisterInfoEndpoints `json:"endpoints"`
}

type SrvRegisterMessage struct {
	Info  SrvRegisterInfo `json:"info"`
	State string          `json:"state"`
}

type WsMethod struct {
	MicroserviceId   string `json:"microserviceId"`
	MicroserviceName string `json:"microserviceName"`
	Name             string `json:"name"`
	Subject          string `json:"subject"`
	FieldKeyName     string `json:"fieldKeyName"`
	FieldKeyValue    string `json:"fieldKeyValue"`
	// TODO: add support of regexp and authenticated roles maybe
	// Domain           string `json:"domain"`
}

var natsConnection *nats.Conn

var methods []WsMethod

func Connect(url string, options nats.Options) error {
	nc, err := nats.Connect(url, nats.UserInfo(options.User, options.Password), nats.PermissionErrOnSubscribe(true))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succesfull connect to NATS!")

	natsConnection = nc
	fmt.Println(methods)
	_, err = natsConnection.Subscribe("$SRV.REGISTER", func(m *nats.Msg) {
		fmt.Println(m.Subject)
		var stringData string = string(m.Data)
		fmt.Println(m.Subject, stringData)
		var jsonData SrvRegisterMessage
		json.Unmarshal(m.Data, &jsonData)
		for _, v := range jsonData.Info.Endpoints {
			fmt.Println(v)
			if len(v.Metadata.Value) <= 0 && len(v.Metadata.ValueRegex) <= 0 {
				break
			}
		}
		fmt.Println("\n", jsonData)
	})

	if err != nil {
		fmt.Println("subscribe on srv info error", err)
	}

	// utils.SetInterval(func() {
	// 	srvInfoInbox := natsConnection.NewInbox()

	// 	srvInfoSubsribe, err := natsConnection.SubscribeSync(srvInfoInbox)
	// 	if err != nil {
	// 		fmt.Println("create sync sub error", err)
	// 	}
	// 	defer srvInfoSubsribe.Unsubscribe()

	// 	err = natsConnection.PublishRequest("$SRV.INFO", srvInfoInbox, []byte(""))
	// 	if err != nil {
	// 		fmt.Println("publish request error", err)
	// 	}

	// 	for {
	// 		msg, err := srvInfoSubsribe.NextMsg(10 * time.Millisecond)
	// 		if err != nil {
	// 			fmt.Println("read srv info messages error", err)
	// 			break
	// 		}
	// 		fmt.Println(string(msg.Data))
	// 	}
	// }, 10*time.Second)

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
