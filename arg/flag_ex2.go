package main

import (
	"flag"
	"fmt"
)

func main() {
	// コマンドライン引数から数字を指定してFizzBuzzのnを指定した値で実行できるようにする
	var n int
	flag.IntVar(&n, "n", 30, "num")
	flag.Parse()
	FizzBuzz(n)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
