package main

import "fmt"

func main() {
	// 以下の switch 文を修正し
	// v が 5 未満の時は v < 5
	// v が 10 未満の時は v < 10
	// v が 20 未満の時は v < 20
	// v が 20 以上の時は何も表示しないようにする
	v := 9
	switch {
	case v < 20:
		fmt.Println("v < 20")
	case v < 10:
		fmt.Println("v < 10")
	case v < 5:
		fmt.Println("v < 5")
	}
}
