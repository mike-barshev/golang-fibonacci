package main

import "sync"

// fibonacci function. Renews 'n-1' and 'n' elements
func fibonacci(x, y *uint64, wg *sync.WaitGroup) {
	*y = *x + *y
	*x = *y - *x
	defer wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	x := new(uint64)
	y := new(uint64)
	*x = 1
	*y = 1
	n := 50 // fibonacci N number
	for i := 1; i < n; i++ {
		wg.Add(1)
		go fibonacci(x, y, wg)
		wg.Wait()
	}
	println(*x)
}