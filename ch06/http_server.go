// 6.5 http server

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Println("Accept %v\n", conn.RemoteAddr())

			// requestの読み込み
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			// responseの書き込み
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}

			response.Write(conn)
			conn.Close()
		}()
	}
}
