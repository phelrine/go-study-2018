package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var showLineNum bool
	flag.BoolVar(&showLineNum, "n", false, "shows line number")
	flag.Parse()
	lines := []string{}
	for _, f := range flag.Args() {
		file, err := os.Open(f)
		if err != nil {
			break
		}
		// ファイルの内容を一行ずつ読む
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			lines = append(lines, string(line))
		}
		// ファイルを閉じる
		file.Close()
	}

	for i, l := range lines {
		if showLineNum {
			fmt.Printf("%d\t%s\n", i+1, l)
		} else {
			fmt.Printf("%s\n", l)
		}
	}
}
