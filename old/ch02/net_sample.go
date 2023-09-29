package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// connがio.Readerインタフェースであることを利用して、レスポンスを画面に出力する
	io.Copy(os.Stdout, conn)
}
