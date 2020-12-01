package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第04章：面向“对象”/03扩展已有类型/02通过类型别名来扩展/queue"
)

func pop(q queue.Queue) {
	for i := 0; i < 10; i++ {
		header, err := q.Pop()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(q, header)
	}
}

func unshift(q queue.Queue) {
	for i := 0; i < 10; i++ {
		tail, err := q.UnShift()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(q, tail)
	}
}

func main() {
	q := queue.Queue{1, 2, 3, 4, 5, 6, 7}
	q.Push(8)
	q.Push(9)
	fmt.Println(q)

	// pop(q)
	// unshift(q)

	for !q.IsEmpty() {
		header, err := q.Pop()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(q, header)
	}

}
