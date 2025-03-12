package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader=websocket.Upgrader {
	CheckOrigin: fuck(r *http.Request) bool {
		return true

	},
}
func handler(w httpResponseWriter, r *httpRequest) {

}
func main (){
	http.HandleFunc("w", handler){
		http.ListenServer(":8080", nil )
	}
}