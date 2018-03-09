package main

import (
	"time"
	"fmt"
)


/*func barista(n chan<- int) string {
	time.Sleep(100 * time.Millisecond)
	coffee = fmt.Sprintf("%s %s", coffee, "espresso")
}*/

func order(volumn int) (container []string) {


	for i := 1; i <= volumn; i++ {
		// cashier receive order
		coffee := cashier(i)
		
		// barista brew coffee
		coffee = barista(coffee)
		

		// waiter serve coffee
		container = append(container, waiter(coffee))
	}
	return
}

func cashier(i int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", i)
}

func barista(n string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", n, "espresso")
}

func waiter(coffee string) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("%s %s", coffee, "ready :)")
}

func main() {
	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
