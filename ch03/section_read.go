// 3.5.1 バイナリ解析 特定部位を切り出す

package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")

	// 14バイト目から7バイトを切り出す
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
