package main

import (
	"fmt"
)

func main() {
	headline := Head("test.txt")
	fmt.Printf("head = %s\n", headline)
	tailline := Tail("test.txt")
	fmt.Printf("tail = %s\n", tailline)
}

// ファイルの先頭行を表示する
func Head(filename string) string {
	return ""
}

// ファイルの最終行を表示する
func Tail(filename string) string {
	return ""
}
