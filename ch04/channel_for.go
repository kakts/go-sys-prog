// 4.2.3 for and channel

package main

import (
	"fmt"
	"math"
)

// primeNumber は素数を返すintのchannelを返す
func primeNumber() chan int {
	result := make(chan int)

	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false

			// 素数チェック 2から√iまでの数で割り切れなければ素数
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}

			// 割り切れない場合は素数としてチャネルに送信
			if !found {
				result <- i
			}
		}
		close(result)
	}()

	return result
}

func main() {
	pn := primeNumber()

	// チャネルから値を受信するたびにループを回す
	for n := range pn {
		fmt.Println(n)
	}
}
