package main

import (
	"errors"
	"fmt"
)

func main() {
	hello()

	hello2("Hello World2")
	hello2("Hello World!")

	fmt.Printf("add(1, 2) = %d\n", add(1, 2))

	ret, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	ret, err = div(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ret = %d\n", ret)

	// ローカル関数の定義
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("f1(10, 20) = %d\n", f1(10, 20))
}

// 引数, 返り値なし関数
func hello() {
	fmt.Println("Hello World")
}

// 引数あり関数
func hello2(text string) {
	fmt.Println(text)
}

// 返り値がある関数
func add(v1, v2 int) int {
	return v1 + v2
}

// 返り値に多値を返す関数
func div(v1, v2 int) (int, error) {
	if v2 == 0 {
		return 0, errors.New("division by zero")
	}
	return v1 / v2, nil
}
