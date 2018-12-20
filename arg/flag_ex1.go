package main

import (
	"fmt"
)

func main() {
	// コマンドライン引数からファイル名を指定して任意のファイルの先頭行を表示できるようにする
	var filename string
	head := Head(filename)
	fmt.Println(head)
}

func Head(filename string) string {
	return ""
}
