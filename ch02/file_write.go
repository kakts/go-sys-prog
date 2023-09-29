// 2.4.1 ファイル出力
package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("os.File example\n"))
}
