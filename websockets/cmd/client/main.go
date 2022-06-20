package main

import (
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: ":5555", Path: "/lower"}

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", resp)

	if c == nil {
		log.Fatalln("invalid websocket nothing to do")
	}

	sampleWords := []string{
		"WoRKER",
		"BooTs",
		"PiCturES",
		"ITEMS",
	}

	// lower messages

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		processed := 0
		for {
			log.Println("starting .... to read here")
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("got this from server %s", message)
			processed++
			if processed == len(sampleWords) {
				break
			}
		}
		cm := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "done")
		c.WriteMessage(
			websocket.CloseMessage, cm,
		)
	}()

	for _, word := range sampleWords {
		err = c.WriteMessage(
			websocket.TextMessage, []byte(word),
		)
	}

	wg.Wait()
	c.Close()
	log.Println("all done")
}
