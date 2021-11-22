package websockets

import (
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

func Echo(ws *websocket.Conn) {
	for {
		var err error
		msg := `Message from backend`
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
		} else {
			fmt.Println("Sending")
		}
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println(reply)
		time.Sleep(5 * time.Millisecond)
	}
}
