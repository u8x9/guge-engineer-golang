package main

import (
	"fmt"
	"math"
)

// 常量
func consts() {
	const filename = "foo.bar"
	const a, b = 3, 4 // 声明常量时，未指明数据类型，所以它的类型是待定的
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 因为 a,b 未指明数据类型，可以不用 float64 进行转换
	fmt.Println(filename, c)
}

// 枚举
func enums() {
	const (
		_ = iota
		cpp
		java
		python
		golang
		rust
	)
	fmt.Println(cpp, java, python, golang, rust)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}
