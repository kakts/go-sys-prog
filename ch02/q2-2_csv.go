// q2.2 csv出力
package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {

	outputFile, err := os.Create("q2-2_csv_output.csv")
	if err != nil {
		panic(err)
	}
	// stdout fileに書き出すWriter作成
	writer := io.MultiWriter(os.Stdout, outputFile)

	csvWriter := csv.NewWriter(writer)
	csvWriter.Write([]string{"name", "age", "address"})
	csvWriter.Write([]string{"george", "20", "Tokyo"})
	csvWriter.Flush()

}
