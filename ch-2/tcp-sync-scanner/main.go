package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Printf("горутина %d начала выполнение\n", p)

		wg.Done()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
		// fmt.Printf("горутина %d начала выполнение\n", i)
	}
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			address := fmt.Sprintf("ratkga.ru:%d", i)
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
	close(ports)
}
