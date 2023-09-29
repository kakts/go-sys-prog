package main

import (
	"fmt"
	"os"
)

func main() {
	// io.Writerインタフェースのfileインスタンス作成
	file, err := os.Create("q2-1-output.txt")
	if err != nil {
		panic(err)
	}
	var name = "John"
	var num = 1
	var flnum = 1.01
	// 第1引数にio.Writerインタフェースを持つインスタンスを渡して、そこに書き出す
	fmt.Fprintf(file, "Hello %v, num: %d, flnum: %f", name, num, flnum)
}
