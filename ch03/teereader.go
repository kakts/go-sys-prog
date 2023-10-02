// 3.7 io.TeeReaderの利用
// io.TeeReaderは読み込まれた内容を別のio.Writerに書き込む

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)

	// データを読み捨てる
	_, _ = io.ReadAll(teeReader)

	// bufferには残っている
	fmt.Println(buffer.String())
}
