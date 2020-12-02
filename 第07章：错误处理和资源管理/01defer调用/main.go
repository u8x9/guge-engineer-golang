package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/u8x9/guge-engineer-golang/第07章：错误处理和资源管理/01defer调用/fib"
)

func tryDefer() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// panic("出错了")
	// fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile() {
	file, err := os.Create("/tmp/fib.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	f := fib.Fib()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(w, f())
	}
}

func main() {
	tryDefer()
	// writeFile()
}
