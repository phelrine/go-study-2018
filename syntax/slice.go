package main

import "fmt"

func main() {
	// slice の宣言
	var s1 []int64

	// 初期値
	fmt.Printf("s1 = %v\n", s1)

	// 代入
	s1 = []int64{1, 2, 3}
	fmt.Printf("s1 = %v\n", s1)

	// スライスの要素数を調べる
	fmt.Printf("len(s1) = %d\n", len(s1))

	// スライスの特定の要素を参照する
	fmt.Printf("s1[2] = %d\n", s1[2])
	// 要素数以上のインデックスを参照するとエラー
	// fmt.Printf("s1[100] = %d\n", s1[100])

	// スライスに要素を追加する
	s2 := append(s1, 4)
	fmt.Printf("s2 = %v\n", s2)

	// スライスを連結する
	s12 := append(s1, s2...)
	fmt.Printf("s12 = %v\n", s12)

	// スライスから部分スライスをとりだす
	s3 := s2[1:3]
	fmt.Printf("s3 = %v\n", s3)

	// 要素数を予約してスライスを作成する
	// 第一引数はスライスの型, 第二引数は要素数, 第三引数は容量
	// 容量はあらかじめ確保しておくメモリ量, 大きいスライスを作成することがわかっている場合は
	// ここの数字をそれに近い数字にすると動作がはやくなる
	s4 := make([]int, 100)
	fmt.Printf("len(s4) = %d, cap(s4) = %d, s4 = %v\n", len(s4), cap(s4), s4)
	s5 := make([]int, 0, 100)
	fmt.Printf("len(s5) = %d, cap(s5) = %d, s5 = %v\n", len(s5), cap(s5), s5)
}
