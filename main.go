package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	file, _ := os.Open("urls.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	os.MkdirAll("images", os.ModePerm)

	for scanner.Scan() {
		url := scanner.Text()

		filepath := filepath.Join("images", filepath.Base(url))
		out, _ := os.Create(filepath)
		defer out.Close()

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		io.Copy(out, resp.Body)
	}
}
