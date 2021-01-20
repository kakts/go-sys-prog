package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	// 文字列を受け取れるが、io.Writerのメソッドではないため、他の構造体では使えない
	buffer.WriteString("bytes.Buffer example\n")

	// 代わりに、io.WriteStringならキャストは不要
	// io.WriteString(&buffer, "bytes.Buffer example\n")
	fmt.Println(buffer.String())
}
