package fib

import (
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func Fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
