package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", i)
			conn, err := net.Dial("tcp", address)
			if err != nil {

				// fmt.Printf("port %d is closed or filtered.\n", i)
				return
			}
			conn.Close()
			fmt.Printf("port %d is open\n", i)
		}(i)
	}
	wg.Wait()
}
