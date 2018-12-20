package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数からファイル名を指定して任意のファイルの先頭行を表示できるようにする
	var filename string
	// エラー処理はスキップ
	filename = os.Args[1]
	head := Head(filename)
	fmt.Println(head)
}

func Head(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}
	return string(line)
}
