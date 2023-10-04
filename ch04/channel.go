// 4.2.1 channel

package main

import "fmt"

func main() {
	fmt.Println("start sub()")

	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()

	// 終了を待つ　値は必要ないので変数で受け取らない
	<-done
	fmt.Println("all tasks are finished")
}
