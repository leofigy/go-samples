package main

import (
	"io"
	"log"
	"net"
	"os"
)

const (
	BUFSIZE = 512
)

func main() {

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":7001"
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Unable to get a new connection", err)
			continue
		}

		go func() {
			if err := Echo(conn); err != nil {
				log.Println(err)
			}
		}()

	}

}

func Echo(conn net.Conn) error {
	defer conn.Close()

	buff := make([]byte, BUFSIZE)
	for {
		size, err := conn.Read(buff[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("got %d bytes: %s", size, string(buff))

		log.Println("echoing ")
		_, err = conn.Write(buff[0:size])

		if err != nil {
			return err
		}
	}
	return nil
}
