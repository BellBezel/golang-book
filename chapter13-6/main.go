package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

var (
	counter int64
	wg 		sync.WaitGroup
)

func main() {
	wg.Add(5)
	go increment(1)
	go increment(2)
	go increment(3)
	go increment(4)
	go increment(5)
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(n int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) 
	}
}

/* OUTPUT :
Final Counter: 10
*/