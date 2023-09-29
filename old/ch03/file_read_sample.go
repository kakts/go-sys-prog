package main

import (
	"io"
	"os"
)

func main() {
	// fileを開く
	file, err := os.Open("file_read.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
