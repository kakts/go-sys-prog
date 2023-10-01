// 2.4.7 JSONを成形してio.Writerに出力

package main

import (
	"encoding/json"
	"io"
	"os"
)

func main() {

	file, err := os.Create("encoded.json")
	if err != nil {
		panic(err)
	}
	// os.Stdoutとfileに書き出す
	multiWriter := io.MultiWriter(os.Stdout, file)

	encoder := json.NewEncoder(multiWriter)

	encoder.SetIndent("", "    ")

	// map[string]stringをエンコードしてos.Stdoutとjsonファイルに書き出す
	encoder.Encode(map[string]string{
		"example": "encodiong/json",
		"hello":   "world",
	})
}
