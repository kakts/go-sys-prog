package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// 複数のio.Writerインタフェースの実体に書き込むためのもの
	// ここではファイルとstdout
	writer := io.MultiWriter(file, os.Stdout)

	io.WriteString(writer, "io.MultiWriter example\n")
}
