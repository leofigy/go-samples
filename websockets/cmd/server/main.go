package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func lower(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	toLow(ws)
	log.Println("final out")
}

func toLow(ws *websocket.Conn) {
	if ws == nil {
		log.Println("invalid web socket ...")
	}

	for {
		mt, message, err := ws.ReadMessage()

		if err != nil {
			log.Println("error msg", err)
			break
		}

		log.Println(mt)

		log.Printf("Message received: %s", message)

		outcome := strings.ToLower(string(message))

		err = ws.WriteMessage(mt, []byte(outcome))

		if err != nil {
			log.Println("error sending response", err)
		}
		log.Printf("Message sent: %s", outcome)
	}
}

func main() {
	http.HandleFunc("/lower", lower)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
