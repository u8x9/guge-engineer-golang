package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello世界和平!"

	fmt.Println("-- string --")
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println()
	fmt.Println("-- byte --")
	for i, b := range []byte(s) {
		fmt.Printf("(%d %X) ", i, b)
	}

	fmt.Println()
	fmt.Println("-- rune --")
	for i, r := range []rune(s) {
		fmt.Printf("(%d %X) ", i, r)
	}

	fmt.Println()

	fmt.Println("len(s) =", len(s))
	fmt.Println("len([]byte(s)) =", len([]byte(s)))
	fmt.Println("len([]rune(s)) =", len([]rune(s)))
	fmt.Println("RuneCountInString =", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}
