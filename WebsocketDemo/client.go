package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)
type Test struct {
	Id string `json:"id"`
	Name string `json:"name"`
}
func main() {
	t := Test{"1651032", "Chair"}
	c, _, _ := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	info, _ := json.Marshal(t)
	defer c.Close()
	_ = c.WriteMessage(websocket.BinaryMessage, info)
	fmt.Print("Finished Sending")
}