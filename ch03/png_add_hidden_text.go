// 3.5.4 PNGに秘密のテキストを埋め込む
// 先頭8byteがシグネチャ
// それ以降はチャンクのブロックで構成される
// チャンクは以下の構成
// 4byte: データ長
// 4byte: チャンクタイプ
// データ長の後にはデータとしてチャンクタイプによって異なるデータが入る
// 4byte: CRC(データ長とチャンクタイプのみを対象にしたCRC-32)

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
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

	// teXtチャンクの場合にのみ中身を表示させる
	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))

	// CRC計算とバッファへの書き込みを同時に行うMultiWriterを用意する
	writer := io.MultiWriter(&buffer, crc)
	// 2バイト目の5ビット目を立てる(小文字にする)とプライベートとなる
	io.WriteString(writer, "teXt")

	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
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

	newFile, err := os.Create("test_secret_text.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)

	// シグネチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])

	// テキストチャンクを追加
	io.Copy(newFile, textChunk("Lambda Note++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}
