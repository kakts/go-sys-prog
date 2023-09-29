package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * テキストファイルから指定したサイズのbyteずつ読み取って、出力用ファイルとstdoutに書き出す。
 * ファイルの終端になるまで続ける
 *
 * TODO 書き出し終了時にデータを重複して書き出してします
 */
func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 4byteのバッファーのメモリ確保
	buffer := make([]byte, 4)

	// 出力用ファイルを作成
	outputFile, err := os.Create("file_read_all_output.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// ファイルとstdoutに出力する用のmultiWriterインスタンス作成
	multiWriter := io.MultiWriter(outputFile, os.Stdout)

	var readByteCount = 0
	for {

		n, err := file.Read(buffer)
		fmt.Printf("read: %d byte read completed.\n", n)
		multiWriter.Write(buffer)
		fmt.Println("")
		if n == 0 {
			fmt.Println("Read 0 byte")
			break
		}
		if err != nil {
			panic(err)
		}

		readByteCount++
		if readByteCount == 10 {
			break
		}
		// 疑問 ファイルの終端をどう判定するか
	}
	fmt.Println("File read finished")
}
