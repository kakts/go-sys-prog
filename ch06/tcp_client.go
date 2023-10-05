// 6.4 tcp client

package main

import "net"

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello"))
}
