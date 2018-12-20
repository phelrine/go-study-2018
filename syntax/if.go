package main

import "fmt"

func main() {
	// 条件分岐
	// 条件式が真の時のみ実行される
	if true {
		fmt.Println("true statement")
	}

	// 条件式にはboolの評価式のみ指定できる
	v := 10
	if v > 5 {
		// 真の時に実行される式
		fmt.Println("v > 5")
	} else {
		// 偽の時に実行される式
		fmt.Println("v <= 5")
	}

	// if 文では条件式以外にも, 評価のために必要な式を実行できる
	v1, v2 := 1, 8
	if v3 := v1 + v2; v3 > 5 {
		fmt.Printf("v3(%d) > 5\n", v3)
	}
	// v3 はスコープ内のみで有効
	// fmt.Printf("v3 = %d\n", v3)

	// 複雑な条件式
	v4 := 15
	if v4 < 5 {
		fmt.Printf("v4(%d) < 5\n", v4)
	} else if v4 < 20 {
		fmt.Printf("v4(%d) < 20\n", v4)
	} else {
		fmt.Printf("v4(%d)\n", v4)
	}

	// switch を利用した定数マッチ
	switch v4 {
	case 5:
		fmt.Printf("v4(%d) == 5\n", v4)
	case 15:
		fmt.Printf("v4(%d) == 15\n", v4)
	default:
		fmt.Printf("v4(%d)\n", v4)
	}

	// switch を利用した条件式マッチ
	switch {
	case v4 < 5:
		fmt.Printf("v4(%d) < 5\n", v4)
	case v4 < 20:
		fmt.Printf("v4(%d) < 20\n", v4)
	default:
		fmt.Printf("v4(%d)\n", v4)
	}
}
