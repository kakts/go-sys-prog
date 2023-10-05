// 6.4 tcp server

package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	_, err = ln.Accept()
	if err != nil {
		panic(err)
	}

}
