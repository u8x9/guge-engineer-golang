package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(id int, w *worker) {
	for n := range w.in {
		fmt.Printf("Worker #%-2d received %3d(%c)\n", id, n, n)
		w.wg.Done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) *worker {
	w := &worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w)
	return w
}

func chanDemon() {
	const num = 10
	var workers [num]*worker
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, w := range workers {
		wg.Add(1)
		w.in <- 'a' + i
	}
	for i, w := range workers {
		wg.Add(1)
		w.in <- 'A' + i
	}
	wg.Wait()

}

func main() {
	chanDemon()
}
