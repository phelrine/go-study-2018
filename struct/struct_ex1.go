package main

import (
	"fmt"
	"math"
)

type Shape interface {
	// 面積を計算するメソッド
	Area() float64
}

type Square struct {
}

func (s Square) Area() float64 {
	return 0.0
}

type Rectangle struct {
}

func (r Rectangle) Area() float64 {
	return 0.0
}

type Triangle struct {
}

func (t Triangle) Area() float64 {
	return 0.0
}

// 1辺の長さxを入力として初期化する
func NewSquare(x float64) *Square {
	return &Square{}
}

// 2辺の長さを入力として初期化する
func NewRectangle(a, b float64) *Rectangle {
	return &Rectangle{}
}

// 2辺(a, b)とその間の角(rad)(ラジアン)を入力として初期化する
func NewTriangle(a, b, rad float64) *Triangle {
	return &Triangle{}
}

func main() {
	var shapes []Shape
	shapes = append(shapes, NewSquare(4))
	shapes = append(shapes, NewRectangle(4, 5))
	shapes = append(shapes, NewTriangle(2, 4, math.Pi/6))

	for _, s := range shapes {
		fmt.Printf("Area() = %f\n", s.Area())
	}
}
