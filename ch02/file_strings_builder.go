// 2.4.4 書かれた内容を記憶しておくバッファ2

package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
