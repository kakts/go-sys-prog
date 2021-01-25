package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		// io.Reader用の5byteバッファ
		buffer := make([]byte, 5)
		// 標準入力から読み込んでbufferに入れる
		size, err := os.Stdin.Read(buffer)

		// EOFの場合終了
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Println("size=%d input='%s'\n", size, string(buffer))

	}
}
