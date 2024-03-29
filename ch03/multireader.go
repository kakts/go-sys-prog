// 3.7 io.MultiReader

package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("-----HEADER-----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	// 全てのreaderをつなげた出力を表示
	io.Copy(os.Stdout, reader)
}
