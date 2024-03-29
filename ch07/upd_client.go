// 7.2.2 udp client

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sending to server")
	// upd packet送信
	_, err = conn.Write([]byte("Hello from client"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")

	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Received: %s\n", string(buffer[:length]))
}
