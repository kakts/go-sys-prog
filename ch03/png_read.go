// 3.5.3 pngファイルの解析
// 先頭8byteがシグネチャ
// それ以降はチャンクのブロックで構成される
// チャンクは以下の構成
// 4byte: データ長
// 4byte: チャンクタイプ
// データ長の後にはデータとしてチャンクタイプによって異なるデータが入る
// 4byte: CRC(データ長とチャンクタイプのみを対象にしたCRC-32)

package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// chunkを読み込む
func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

// readChunks pngの複数のチャンクを読み込む
func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 最初の8byteはスキップ
	file.Seek(8, 0)

	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}

		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所のため
		// チャンク名(4byte) + データ長 + CRC(4byte)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := readChunks(file)

	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
