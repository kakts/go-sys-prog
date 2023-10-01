// 3.4.1 stdin

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		// reader用に5byteのバッファを用意
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size = %d, input ='%s'\n", size, string(buffer))
	}
}
