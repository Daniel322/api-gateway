package natsadapter

import (
	"fmt"
	"time"

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

	_, err = natsConnection.Subscribe("$SRV.>", func(m *nats.Msg) {
		fmt.Println(m.Subject)
		var stringData string = string(m.Data)
		fmt.Println(m.Subject, stringData)
		// var jsonData map[string]map[string]any
		// json.Unmarshal(m.Data, &jsonData)
		// fmt.Println(m.Subject, "jsonData", jsonData)
	})

	if err != nil {
		fmt.Println("subscribe on srv info error", err)
	}

	srvInfoInbox := natsConnection.NewInbox()

	srvInfoSubsribe, err := natsConnection.SubscribeSync(srvInfoInbox)
	if err != nil {
		fmt.Println("create sync sub error", err)
		return err
	}
	defer srvInfoSubsribe.Unsubscribe() // Гарантированная отписка

	err = natsConnection.PublishRequest("$SRV.INFO", srvInfoInbox, []byte(""))
	if err != nil {
		fmt.Println("publish request error", err)
		return err
	}

	for {
		msg, err := srvInfoSubsribe.NextMsg(10 * time.Millisecond)
		if err != nil {
			fmt.Println("read srv info messages error", err)
			break
		}
		fmt.Println(string(msg.Data))
	}

	// TODO: make such inbox with SRV.PING

	return err
}

func Disconnect() {
	natsConnection.Close()
	fmt.Println("succesfull disconect from NATS")
}
