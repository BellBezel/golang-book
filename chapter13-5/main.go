package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter int
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
		value := counter
		runtime.Gosched()   // <-- ทำให้เกิด Data less
		value++
		counter = value
	}
}

/* Without 'runtime.Gosched()' OUTPUT : 
Final Counter: 10
Final Counter: 10
Final Counter: 10
Final Counter: 10
Final Counter: 10

With 'runtime.Gosched()' OUTPUT :
Final Counter: 2
Final Counter: 2
Final Counter: 2
Final Counter: 8
*/