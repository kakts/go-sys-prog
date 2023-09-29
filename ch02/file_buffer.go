// 2.4.3 書かれた内容を記憶しておくバッファ

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
