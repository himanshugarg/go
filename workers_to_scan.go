package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for i := range ports {
		fmt.Println(i)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ports := make(chan int, 100)

	// workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	// ports
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
