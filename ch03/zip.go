// 3.9 Q3.3 zipファイルの書き込み

package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	// zipファイルを作成
	zipFile, err := os.Create("output_zip.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	// zipファイル用のio.Writerを取得
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// ファイルを追加
	err = addToZip("zip_sample1.txt", zipWriter)
	if err != nil {
		panic(err)
	}

	err = addToZip("zip_sample1.txt", zipWriter)
	if err != nil {
		panic(err)
	}
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 指定したファイル名でzipファイルに書き込みする用のio.Writerを取得
	writer, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	// ファイルの内容をzipに書き込み
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}
