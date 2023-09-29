package main

// 3.5.3 PNGファイルの分析

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes) \n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 最初の8byte(PNGのシグニチャ)を飛ばす
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		// 次のブロックを読み取る
		// 長さ(4byte)種類(4byte)データ(長さで指定した数)CRC(4byte)
		// データの長さと、それ以外の4byte*3で12byteを足す
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
	}
	return chunks
}

func main() {
	file, err := os.Open("test-png.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
