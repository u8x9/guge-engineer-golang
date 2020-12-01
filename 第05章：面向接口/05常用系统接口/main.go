package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// fmt.Stringer
func (p *person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}

// io.Reader
func (p *person) Read(b []byte) (int, error) {
	// WRONG!
	copy(b, []byte(p.name))
	return len(b), nil
}

// io.Writer
func (p *person) Write(b []byte) (int, error) {
	p.name = string(b)
	return len(b), nil
}

func main() {
	p := &person{"张三", 123}
	fmt.Println(p)

	fmt.Fprint(p, "李四")
	fmt.Println(p)

	// ioutil.ReadAll(p)
}
