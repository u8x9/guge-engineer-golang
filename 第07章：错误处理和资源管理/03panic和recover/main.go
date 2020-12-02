package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	// panic(errors.New("this is an error"))

	// x, y := 123, 0
	// fmt.Println(x / y)

	panic(123)
}

func main() {
	tryRecover()
}
