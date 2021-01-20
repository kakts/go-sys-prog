package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write()が受け取るのは文字列でなくバイト列なので、変換を行ってからWrite()に渡す
	file.Write([]byte("os.File example\n"))
}
