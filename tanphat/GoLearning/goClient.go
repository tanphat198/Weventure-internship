package GoLearning

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)


type Message struct {
	Person string `json:"person"`
	Content string `json:"content"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	c, _, _ := websocket.DefaultDialer.Dial("localhost:8080", nil)
	message := &Message{
		Person:"Phat",
		Content:"EH! may",
	}
	bytesMessage,_ :=json.Marshal(message)

	err := c.WriteMessage(websocket.TextMessage,bytesMessage)

	if err != nil {
		log.Println("write close:", err)
		return
	}
	fmt.Println("sended")

}



func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

