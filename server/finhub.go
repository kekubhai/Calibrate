package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func connectToFinnhub() {
	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token=cv96jhpr01qjq626urggcv96jhpr01qjq626urh0", nil)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	symbols := []string{"AAPL", "AMZN", "BINANCE:BTCUSDT", "IC MARKETS:1"}
	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msg)
	}

	var msg interface{}
	for {
		err := w.ReadJSON(&msg);
		if err != nil {
			panic(err)
		}
		fmt.Println("Message from server ", msg)
	}
}

