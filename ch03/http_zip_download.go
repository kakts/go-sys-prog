// 3.9 Q3.4 net/httpのRequestを利用したzipダウンロード

package main

import (
	"io"
	"net/http"
	"os"
)

// zipファイルをダウンロードする
func downloadZip(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	file, err := os.Open("output_zip.zip")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// responseとしてzipファイルを返す
	io.Copy(w, file)
}

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

	http.HandleFunc("/", downloadZip)
	http.ListenAndServe(":8080", nil)
}
