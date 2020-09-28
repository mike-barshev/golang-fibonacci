package main

import (
	"sync"
)

// fibonacci function
func fibonacci(x, y *int, wg *sync.WaitGroup) {
	*y = *x + *y
	*x = *y - *x
	defer wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	x := new(int)
	y := new(int)
	*x = 1
	*y = 1
	n := 10 // Число n Фибоначчи
	for i := 1; i < n; i++ {
		wg.Add(1)
		go fibonacci(x, y, wg)
		wg.Wait()
	}
	println(*x)
}
