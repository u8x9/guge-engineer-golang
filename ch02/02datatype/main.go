package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// euler 欧拉公式（复数的使用）
func euler() {
	// c := 3 + 4i
	// result := cmplx.Abs(c)
	// fmt.Println(result)
	r := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f\n", r)
}

// triangle 类型转换是强制的
func triangle() {
	var a, b int = 3, 4
	var c int
	// c = math.Sqrt(a*a + b*b)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
