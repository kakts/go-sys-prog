// 3.4.3 ネットワーク通信の読み込み

package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))

	// レスポンス結果をstdoutに出力
	io.Copy(os.Stdout, conn)
}
