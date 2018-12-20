package main

import "fmt"

func main() {
	FizzBuzz(100)
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
