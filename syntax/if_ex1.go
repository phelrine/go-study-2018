package main

import "fmt"

func main() {
	v := 10
	// 以下の条件式を修正し v が奇数の時は奇数, 偶数の時は偶数と表示されるように修正する
	if v%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}
