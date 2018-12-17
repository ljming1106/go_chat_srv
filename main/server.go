package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn fail ...,", err)
		}
		fmt.Println("connect client successfully")
		var msg string
		for {
			msg = ""
			data := make([]byte, 255)
			msg_read, err := conn.Read(data)
			if msg_read == 0 || err != nil {
				break
			}

			msg_read_str := string(data[0:msg_read])
			if msg_read_str == "close" {
				conn.Write([]byte("close"))
				break
			}

			fmt.Println("client say: ", msg_read_str)

			fmt.Printf("say to client: ")
			fmt.Scan(&msg)
			conn.Write([]byte(msg))
		}
		fmt.Println("client Close")
		conn.Close()
	}
}
