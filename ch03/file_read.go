// 3.4.2 ファイル入力

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 実行中のファイル名を取得
	dirName, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dirName = %s\n", dirName)

	file, err := os.Open(dirName + "/ch03/file_read.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// stdoutにファイルの内容を全て出力
	io.Copy(os.Stdout, file)

}
