package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var p []byte
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	_, p, _ = ws.ReadMessage()
	_ = ws.WriteMessage(websocket.BinaryMessage, p )
	fmt.Println("Successfully Received")

}
func python(w http.ResponseWriter, r *http.Request){
	ws, _ := upgrader.Upgrade(w, r, nil)
	_ = ws.WriteMessage(websocket.BinaryMessage, p)
	fmt.Println("Successfully Sent")
}


func main() {
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/python", python)
	log.Fatal(http.ListenAndServe(":8080", nil))
}