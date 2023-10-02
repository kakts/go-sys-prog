// 3.6.1 テキスト解析用io.Reader実装
// 改善版 任意の文字でテキスト分割する
// bufio.Scannerを使うことにより、終端文字を指定する必要がなくなる ただしbuf.io.Readerと違い分割文字が削除されている

package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
