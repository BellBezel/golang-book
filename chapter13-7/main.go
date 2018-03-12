package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

var (
	counter int64
	wg 		sync.WaitGroup
	mu		sync.Mutex
)

func main() {
	wgmun := 16
	wg.Add(wgmun)
	for i:=1;i<=wgmun;i++ {
		go increment(i)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(n int) {
	defer wg.Done()
	mu.Lock()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) 
	}
	mu.Unlock()
}

/* OUTPUT :
Final Counter: 32
*/