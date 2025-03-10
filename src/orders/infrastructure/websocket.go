package infrastructure

import (
	"log"
	"apiConsumer/src/orders/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func runWebsocket(order *domain.Order) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5000/ws", nil)
	if err != nil {
		log.Println("Error connecting to WebSocket:", err)
		return
	}
	defer conn.Close()

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("Error encoding order to JSON:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, orderJSON); err != nil {
		log.Println("Error sending order:", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Println("Received message from server:", string(message))
	}
}