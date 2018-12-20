package main

import "fmt"

func main() {
	fmt.Printf("Sum(10) = %d\n", Sum(10))
}

// 0 から n までの数字の和を計算する関数を作成する
func Sum(n int) int {
	ret := 0
	for i := 0; i <= n; i++ {
		ret += i
	}
	return ret
}
