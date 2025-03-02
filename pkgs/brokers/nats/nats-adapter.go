package natsadapter

import (
	"encoding/json"
	"fmt"
	"time"
	"websocket-gateway/pkgs/utils"

	nats "github.com/nats-io/nats.go"
)

type Endpoint struct {
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	Metadata any    `json:"metadata"`
}
type Method struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Endpoints []interface{} `json:"endpoints"`
}

var natsConnection *nats.Conn

var methods []Method

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
		var jsonData map[string]map[string]any
		json.Unmarshal(m.Data, &jsonData)
		methods = append(
			methods,
			Method{
				Id:        jsonData["info"]["id"].(string),
				Name:      jsonData["info"]["name"].(string),
				Endpoints: jsonData["info"]["endpoints"].([]interface{}),
			},
		)
		fmt.Println("\n", methods)
	})

	if err != nil {
		fmt.Println("subscribe on srv info error", err)
	}

	utils.SetInterval(func() {
		srvInfoInbox := natsConnection.NewInbox()

		srvInfoSubsribe, err := natsConnection.SubscribeSync(srvInfoInbox)
		if err != nil {
			fmt.Println("create sync sub error", err)
		}
		defer srvInfoSubsribe.Unsubscribe()

		err = natsConnection.PublishRequest("$SRV.INFO", srvInfoInbox, []byte(""))
		if err != nil {
			fmt.Println("publish request error", err)
		}

		for {
			msg, err := srvInfoSubsribe.NextMsg(10 * time.Millisecond)
			if err != nil {
				fmt.Println("read srv info messages error", err)
				break
			}
			fmt.Println(string(msg.Data))
		}
	}, 10*time.Second)

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
