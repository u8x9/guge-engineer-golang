package main

import (
	"fmt"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Printf("%d %q\n", a, s)
}
func variableTypeDeduction() {
	var a, s = 3, "abc"
	fmt.Printf("%T, %T\n", a, s)
}

func variableShorter() {
	a, s := 3, "abc"
	fmt.Printf("%T, %T\n", a, s)
}
func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	fmt.Println("Hello, 世界!")
}
