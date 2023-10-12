// 9.2.1 ファイルへの追記

package main

import (
	"io"
	"os"
)

// append ファイルを追記モードで開き、書き込む
func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func main() {
	append()
}
