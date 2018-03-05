package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup	//จะเป็นตัวจำ queue ว่าใครมาก่อน-หลัง
	wg.Add(2)
	
	for i:=0; i < 2; i++ {
		go func(n int) {
			defer wg.Done()		// ทำงานเมื่อจบการรันทั้งหมดแล้ว
			for i := 0; i < 10; i++ {
				fmt.Println(n, ":", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}

/* OUTPUT :
0 : 0
0 : 1
0 : 2
0 : 3
0 : 4
0 : 5
0 : 6
0 : 7
0 : 8
0 : 9
1 : 0
1 : 1
1 : 2
1 : 3
1 : 4
1 : 5
1 : 6
1 : 7
1 : 8
1 : 9
Finished
*/