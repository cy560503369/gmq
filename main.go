package main

import (
	"fmt"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
	"net"
	"os"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		go func(a int) {
			conn, err := net.Dial("tcp", "139.198.123.37:1883")
			if err != nil {
				fmt.Fprint(os.Stderr, "dial: ", err)
				return
			}

			cc := mqtt.NewClientConn(conn)
			cc.Dump = false

			if err := cc.Connect("", ""); err != nil {
				fmt.Fprint(os.Stderr, "connect: %v\n", err)
			}

			str := fmt.Sprintf("000000000000000%d", a)
			cc.ClientId = str
			fmt.Println("connect with client id", cc.ClientId)

			for {
				cc.Publish(&proto.Publish{
					Header:    proto.Header{},
					TopicName: "eems/td/command/10180501100335",
					Payload:   proto.BytesPayload([]byte("121321143143315")),
				})

				time.Sleep(2 * time.Second)
			}

			cc.Disconnect()
		}(i + 1)
	}
	fmt.Println("end")
	for {
		time.Sleep(1 * time.Hour)
	}
}
