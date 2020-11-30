package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) (result string) {
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}
func forever() {
	for {
		fmt.Println("foobar")
	}
}

// 逐行打印文件内容
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(7128713123),
		convertToBin(0),
	)

	printFile("./main.go")

	// forever()
}
