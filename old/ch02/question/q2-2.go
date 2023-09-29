package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	// 書き出すファイル作成
	file, err := os.Create("q2-2-output.csv")
	if err != nil {
		panic(err)
	}

	// csvファイルとstdoutに書き出すwriterインスタンスを作成
	writer := io.MultiWriter(file, os.Stdout)

	csvWriter := csv.NewWriter(writer)
	csvWriter.Write([]string{"test,a,1", "test,b,2"})
	// Flushによってバッファデータを指定したio.Writerに書き出す。
	csvWriter.Flush()
}
