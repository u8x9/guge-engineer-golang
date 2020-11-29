package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "main.go"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", contents)
}
