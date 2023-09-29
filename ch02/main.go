// 2.2 Goのインタフェース

package main

import "fmt"

// Talker は 会話することができる型に実装される
type Talker interface {
	// Talk は 会話を行う
	Talk()
}

// Greeter は 会話を行う主体を表す
type Greeter struct {
	// name は話者の名前です
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"George"}
	talker.Talk()
}
