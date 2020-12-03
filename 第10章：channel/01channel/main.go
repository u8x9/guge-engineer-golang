package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func worker(c <-chan int) {
	// for {
	// 	n := <-c
	// 	fmt.Println(n)
	// }

	// for {
	// 	if n, ok := <-c; ok { // close 之后，ok 为 false
	// 		fmt.Println(n)
	// 	} else {
	// 		break
	// 	}
	// }

	for n := range c { // close 之后，循环结束
		fmt.Println(n)
	}
}

func worker1(id int, c <-chan int) {
	for {
		fmt.Printf("Worker #%d received %c\n", id, <-c)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker1(id, c)
	return c
}
func task(c chan<- int) {
	c <- 1
	c <- 2
}
func chanDemo2() {
	c := make(chan int)
	go worker(c)
	go task(c)
	time.Sleep(time.Millisecond)
}
func chanDemo3() {
	const num = 10
	var chs [num]chan int
	for i := 0; i < num; i++ {
		chs[i] = make(chan int)
		go worker1(i, chs[i])
	}
	for i := 0; i < num; i++ {
		chs[i] <- 'a' + i
	}
	for i := 0; i < num; i++ {
		chs[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func chanDemo4() {
	const num = 10
	var chs [num]chan<- int
	for i := 0; i < num; i++ {
		chs[i] = createWorker(i)
	}
	for i := 0; i < num; i++ {
		chs[i] <- 'a' + i
	}
	for i := 0; i < num; i++ {
		chs[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	// 由发送方 close
	close(c) // close之后，接收方依然会收到 chan 的零值
	time.Sleep(time.Millisecond)
}
func main() {
	// chanDemo()
	// chanDemo2()
	// chanDemo3()
	// chanDemo4()
	// bufferedChannel()
	channelClose()
}
