// 2.4.7 フォーマットしたデータをio.Writerに出力

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
