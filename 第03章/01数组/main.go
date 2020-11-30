package main

import (
	"fmt"
)

// ⚠️ 数组是值类型

func main() {
	var a1 [5]int
	a2 := [3]int{1, 2, 3}
	a3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(a1, a2, a3)
	fmt.Println(grid)
}
