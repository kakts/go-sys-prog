package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	encoder := json.NewEncoder(io.MultiWriter(gzipWriter, os.Stdout))
	encoder.SetIndent("", "   ")
	encoder.Encode(source)
	// io.MultiWriter(gzipされたjson)
	// gzip後: file
	// gzip前: stdout
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
