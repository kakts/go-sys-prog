package main

import (
	"io"
	"net"
	"os"
)

func main() {
	// net.Conn型のインスタンス取得
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	// stdoutにコピーすることでデータを一括で読み込み
	// この場合読み込まれるのは生のHTTP通信内容そのもの
	io.Copy(os.Stdout, conn)
}
