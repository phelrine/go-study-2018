package main

import "fmt"

func main() {
	// 繰り返し構文
	for i := 0; i < 10; i++ {
		// 10回実行される
		fmt.Printf("i = %d\n", i)
	}

	// 条件文のみでも繰り返し可能
	j := 1
	for j < 20 {
		fmt.Printf("j = %d\n", j)
		j *= 2
	}

	// スライスの要素を列挙する
	// i が index, e が要素として代入される
	s := []int{1, 2, 3, 4, 5}
	for i, e := range s {
		fmt.Printf("i = %d, e = %d\n", i, e)
	}

	// 文字列の要素を列挙する
	str := "Hello World"
	for i, c := range str {
		fmt.Printf("i = %d, c = %c\n", i, c)
	}
}
