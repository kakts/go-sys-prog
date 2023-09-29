// 2.4.5 インターネットアクセスの送信
// io.Writerによるインターネットアクセス
// net.Connを利用する
// net.Connはio.Writer, ioReaderを実装している

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

	// 送信
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")

	file, err := os.Create("httpout.txt")
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)

	// サーバからのレスポンスをStdoutとhttpout.txtに出力
	io.Copy(writer, conn)

}
