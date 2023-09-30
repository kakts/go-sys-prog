// 2.9
// Q2.1

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("q2-1_output.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "int: %d, 文字列: %s, float: %f", 10, "Hello", "1.11")
}
