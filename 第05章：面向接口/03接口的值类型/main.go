package main

import (
	"fmt"
	"time"

	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver/mock"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver/real"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/03接口的值类型/queue"
)

func whatIsInterface(r retriver.Retriver) {
	fmt.Printf("%T %#v\n", r, r)
}

func showInterfaceType(r retriver.Retriver) {
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("mock: ", v.Content)
	case *real.Retriver:
		fmt.Println("real: ", v.UserAgent)
	default:
		fmt.Println("unknown")
	}
}

func typeAssertion(r retriver.Retriver) {
	rr, ok := r.(*real.Retriver)
	if !ok {
		fmt.Println("is not *real.Retriver")
		return
	}
	fmt.Println("type assertion real: ", rr.UserAgent, rr.Timeout)
}

func main() {
	var r retriver.Retriver
	whatIsInterface(r)
	showInterfaceType(r)
	typeAssertion(r)
	fmt.Println()

	r = &mock.Retriver{Content: "foobar"}
	whatIsInterface(r)
	showInterfaceType(r)
	typeAssertion(r)
	fmt.Println()

	r = &real.Retriver{
		UserAgent: "u8x9/go http client",
		Timeout:   time.Minute,
	}
	whatIsInterface(r)
	showInterfaceType(r)
	typeAssertion(r)

	var q queue.Queue
	for i := 0; i < 10; i++ {
		q.Push(i + 1)
		q.Push(fmt.Sprintf("%c", 65+i))
	}
	fmt.Println(q)
	for i := 0; i < 20; i++ {
		// v := q.Pop()
		v := q.UnShift()
		fmt.Print(q, v)
		switch v.(type) {
		case int:
			fmt.Print(" [INT]")
		case string:
			fmt.Print(" [STRING]")
		}
		fmt.Println()
	}
}
