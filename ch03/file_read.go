package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("hello.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 4)

	// fileから4byte読み込み
	// ファイルの先頭から4byteのみで、全体を読み込まない
	n, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file read. %d", n)

	fmt.Println("Write buffer to hello_out.txt")
	outputFile, err := os.Create("hello_out.txt")
	defer outputFile.Close()
	if err != nil {
		panic(err)
	}
	// 4byte分のバッファーをアウトプットに書き出し
	outputFile.Write(buffer)
	fmt.Println("Write buffer done")
}
