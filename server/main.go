package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("Recieved : %s\n", message)
		res := fmt.Sprintf("%s, at: %s", string(message), time.Now().String())
		if err := conn.WriteMessage(messageType, []byte(res)); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("Server started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
