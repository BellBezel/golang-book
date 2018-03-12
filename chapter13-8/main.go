package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a, b value
	var wg sync.WaitGroup
	wg.Add(2)
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)
	wg.Wait()
}

type value struct {
	mu    sync.Mutex
	value int
}

func printSum(a, b *value, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock() // introduce deadlock
	time.Sleep(2 * time.Second)
	b.mu.Lock()
	defer b.mu.Unlock() // introduce deadlock
	fmt.Printf("sum=%v\n", a.value+b.value)
}

/* OUTPUT :
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc0420100ac)
        C:/Go/src/runtime/sema.go:56 +0x40
sync.(*WaitGroup).Wait(0xc0420100a0)
        C:/Go/src/sync/waitgroup.go:131 +0x79
main.main()
        C:/Users/BELL/Desktop/src/dojo/golang-book/chapter13-8/main.go:15 +0x102

goroutine 5 [semacquire]:
sync.runtime_SemacquireMutex(0xc042010094, 0x0)
        C:/Go/src/runtime/sema.go:71 +0x44
sync.(*Mutex).Lock(0xc042010090)
        C:/Go/src/sync/mutex.go:134 +0xf5
main.printSum(0xc042010080, 0xc042010090, 0xc0420100a0)
        C:/Users/BELL/Desktop/src/dojo/golang-book/chapter13-8/main.go:28 +0xa9
created by main.main
        C:/Users/BELL/Desktop/src/dojo/golang-book/chapter13-8/main.go:13 +0xbe

goroutine 6 [semacquire]:
sync.runtime_SemacquireMutex(0xc042010084, 0x0)
        C:/Go/src/runtime/sema.go:71 +0x44
sync.(*Mutex).Lock(0xc042010080)
        C:/Go/src/sync/mutex.go:134 +0xf5
main.printSum(0xc042010090, 0xc042010080, 0xc0420100a0)
        C:/Users/BELL/Desktop/src/dojo/golang-book/chapter13-8/main.go:28 +0xa9
created by main.main
        C:/Users/BELL/Desktop/src/dojo/golang-book/chapter13-8/main.go:14 +0xf4
exit status 2
*/
