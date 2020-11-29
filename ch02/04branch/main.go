package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) (s string) {
	switch {
	case score <60:
		s="F"
	case score <70:
		s="D"
	case score<80:
		s="C"
	case score <90:
		s="B"
	case score<=100:
		s="A"
	case score<0||score>10:
		panic("非法数据")
	}

	return
}
func main() {
	const filename = "main.go"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", contents)

	s:=grade(91)
	fmt.Println(s)
}
