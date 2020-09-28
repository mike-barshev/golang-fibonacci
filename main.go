package main

import "sync"

func unaryExpression(x *int, wg *sync.WaitGroup) {
	*x++
	defer wg.Done()
}

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
	i := new(int)
	*x = 1
	*y = 1
	*i = 1

	n := 50 // fibonacci N number

	for *i < n {
		wg.Add(2)
		go unaryExpression(i, wg)
		go fibonacci(x, y, wg)
		wg.Wait()
	}
	println(*x)
}