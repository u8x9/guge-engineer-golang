package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// “正统” 的函数式编程实现 adder
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func main() {
	// a := adder()
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(a(i))
	// }

	a := adder2(0)
	for i := 1; i <= 100; i++ {
		var s int
		s, a = a(i)
		fmt.Println(s)
	}
}
