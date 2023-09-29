// 2.4.7 net/httpのRequestを利用した例

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダー")

	outputFile, err := os.Create("http_request_data.txt")
	if err != nil {
		panic(err)
	}

	// os.Stdoutとfileに書き出す
	writer := io.MultiWriter(os.Stdout, outputFile)
	// リクエストデータをos.Stdoutとfileに出力する
	request.Write(writer)
}
