package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/websocket"
)

type Message struct {
	MessageType int
	Data        []byte
}

func main() {
	u := url.URL{Scheme: "ws", Host: "LocalHost:3000", Path: "/ws"}
	fmt.Printf("Connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial", err)
	}
	defer conn.Close()

	send := make(chan Message)
	done := make(chan struct{})
	//Goroutines for reading Messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("Recieved %s\n", message)
		}
	}()

	///Go routines for sending Messages

}
