// 6.4 tcp server

package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// リクエストを受け取って1回で終了しないために、Accept()を何度も呼ぶ
	for {
		_, err := ln.Accept()
		if err != nil {
			// handle error
		}

		// 1リクエスト処理中に他のリクエストのAccept()が実行できるように、goroutineを使って非同期にレスポンス処理する
		go func() {
			// connを使った読み書き

		}()

	}
}
