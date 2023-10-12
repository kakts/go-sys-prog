package main

import (
	"fmt"
	"io"
	"os"
)

const targetFile = "textfile.txt"

// open ファイルの新規作成
func open() {
	file, err := os.Create(targetFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ファイルに文字列を書き込む
	io.WriteString(file, "new file content\n")
}

// read ファイルの読み込み
func read() {
	file, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Read file:")
	// ファイルの内容をstdoutに出力
	io.Copy(os.Stdout, file)
}

func main() {
	open()
	read()
}
