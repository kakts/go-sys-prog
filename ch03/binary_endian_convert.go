// 3.5.2 エンディアン変換
// ビッグエンディアンをリトルエンディアンに変換

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32bit big endian data
	data := []byte{0x0, 0x0, 0x27, 0x10}

	var i int32
	// エンディアン変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)

	fmt.Printf("data: %d\n", i)
}
