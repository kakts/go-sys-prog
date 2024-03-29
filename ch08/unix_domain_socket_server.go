// 8.2.1 unix domain socket http server

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// unix domain socket用のファイルのパスを作成
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
	os.Remove(path)

	// unix domain socketのリスナーを生成
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("server is running at " + path)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("Accept: %v\n", conn.RemoteAddr())
			// リクエストの読み込み
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			// レスポンスの書き込み
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 1,
				Body:       io.NopCloser(strings.NewReader("Hello world\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
