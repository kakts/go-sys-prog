// 7.2.1 UDP サーバ

package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Server is running at localhost:8888")

	// udp socket listen
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		// udp packet受信し、bufferに格納
		// ReadFromによって、受信するのと同時に、接続してきた相手のアドレス情報を受け取れる
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		// udp packet送信
		_, err = conn.WriteTo([]byte("hello from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}

}
