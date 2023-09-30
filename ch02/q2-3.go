// q2-3

// サーバ起動後、clientで
// gzipされたjsonデータをファイルに出力
//   curl -s -v --compressed --raw -o result.gz localhost:8080を実行し
// 出力されたファイルのgzipを解凍して中身を確認する
//  gunzip result.gz

package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	// jsonに変換
	jsonBytes, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	// gzip圧縮を行うのと、stdoutにも圧縮前のデータを出力する
	writer := io.MultiWriter(gzipWriter, os.Stdout)
	io.WriteString(writer, string(jsonBytes))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
