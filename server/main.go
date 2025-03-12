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
conn , err:=upgrader.Upgrade(w , r , nil)
if err!=nil {
	return (err)

}
defer conn.Close()
}
func main (){
	http.HandleFunc("w", handler){
		http.ListenServer(":8080", nil )
	}
}