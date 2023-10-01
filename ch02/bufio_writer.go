// 2.4.6 bufio.Writerを利用した例
// Flush()を呼び出すと、後続のio.Writerに書き込まれる
// 呼び出さないと、バッファデータを保持した状態で消滅する

package main

import (
	"bufio"
	"os"
)

func main() {
	// 引数に入れたio.Writerに対して書き出す
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")

	// バッファの内容を出力する
	buffer.Flush()

	buffer.WriteString("example\n")
	buffer.Flush()
}
