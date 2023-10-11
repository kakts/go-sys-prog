// 8.2.3 unix domain socket http client

package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	// unix domain socketに接続
	conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
	if err != nil {
		panic(err)
	}

	// 第2引数でホスト名を指定しているが、接続先をファイル名で指定しているため、この情報はヘッダーのみで使われる
	request, err := http.NewRequest("get", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
