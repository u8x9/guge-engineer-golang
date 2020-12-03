package main

import (
	"fmt"
)

type notify chan struct{}
type worker struct {
	in   chan int
	done notify
}

func doWorker(id int, w *worker) {
	for n := range w.in {
		fmt.Printf("Worker #%-2d received %3d(%c)\n", id, n, n)
		go func() {
			w.done <- struct{}{}
		}()
	}
}

func createWorker(id int) *worker {
	w := &worker{
		in:   make(chan int),
		done: make(notify),
	}
	go doWorker(id, w)
	return w
}

func chanDemon() {
	const num = 10
	var workers [num]*worker
	for i := 0; i < num; i++ {
		workers[i] = createWorker(i)
	}

	// 问题：造成顺序执行
	// for i := 0; i < num; i++ {
	// 	workers[i].in <- 'a' + i
	// 	<-workers[i].done
	// }
	// for i := 0; i < num; i++ {
	// 	workers[i].in <- 'A' + i
	// 	<-workers[i].done
	// }

	for i, w := range workers {
		w.in <- 'a' + i
	}
	for i, w := range workers {
		w.in <- 'A' + i
	}
	for _, w := range workers {
		<-w.done
		<-w.done
	}
}

func main() {
	chanDemon()
}
