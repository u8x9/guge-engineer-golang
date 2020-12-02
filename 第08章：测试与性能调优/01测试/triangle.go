package main

import (
	"fmt"
	"math"
)

func calcTriangle(a, b int) (c int) {
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}
