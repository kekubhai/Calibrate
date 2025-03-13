package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

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
		defer close(done)
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
	go func() {
		for {
			select {
			case msg := <-send:
				//write in the websocket
				err := conn.WriteMessage(msg.MessageType, msg.Data)
				if err != nil {
					log.Println("write ", err)
					return
				}
			case <-done:
				return
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type Something Here ..... \n")
	for scanner.Scan() {
		text := scanner.Text()

		send <- Message{websocket.TextMessage, []byte(text)}
	}
	if err := scanner.Err(); err != nil {
		log.Println(scanner.Err())
	}
}
