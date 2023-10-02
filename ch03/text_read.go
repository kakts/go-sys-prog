// 3.6.1 テキスト解析用io.Reader実装
// 任意の文字でテキスト分割する

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行目
2行目
3行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))

	for {
		// sourceから改行が見つかるまで読み込む
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
