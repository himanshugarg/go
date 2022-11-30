package main

import (
	"fmt"
	"net"
	"sync"
	"os"
)

func worker(host string, ports chan int, wg *sync.WaitGroup) {
	for i := range ports {
		address := fmt.Sprintf("%s:%d", host, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return
		}
		conn.Close()
		fmt.Printf("%s open\n", address)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	var ports = make(chan int, 100)

  	host := "scanme.nmap.org"
  	if len(os.Args) == 2 {
  	  	host = os.Args[1]
  	}

	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, &wg)
	}

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
}
