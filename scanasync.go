package main

import (
	"fmt"
	"net"
	"sync"
	"os"
)

func main() {
	var wg sync.WaitGroup
  	host := "scanme.nmap.org"
  	if len(os.Args) == 2 {
  	  host = os.Args[1]
  	}
	for i := 1; i <= 65535; i++ {
		wg.Add(1);

		go func(i int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s open\n", address)
		}(i)
	}
	wg.Wait()
}
