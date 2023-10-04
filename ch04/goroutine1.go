// 4.1 goroutine

package main

import (
	"fmt"
	"time"
)

// sub 新しく作られるgoroutineで実行される
func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("Start sub()")

	// goroutineを使って関数実行
	go sub()
	time.Sleep(2 * time.Second)
	fmt.Println("Finish main()")
}
