package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	ReadFile("test.txt")
	ReadFileLine("test.txt")
}

func ReadFile(filename string) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルの内容全てを読む
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(data))

	// ファイルを閉じる
	file.Close()
}

func ReadFileLine(filename string) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルの内容を一行ずつ読む
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// ファイルを閉じる
	file.Close()
}
