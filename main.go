package main

import "sync"

func fibonacci(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	wg := new(sync.WaitGroup)
	x := new(uint64)
	y := new(uint64)
	i := new(int)
	*x = 1
	*y = 1
	*i = 1
	for *i < n {
		wg.Add(2)
		go fibonacciRenews(x, y, wg)
		go unaryExpression(i, wg)
		wg.Wait()
	}
	return *x
}

func unaryExpression(x *int, wg *sync.WaitGroup) {
	*x++
	defer wg.Done()
}

func fibonacciRenews(x, y *uint64, wg *sync.WaitGroup) {
	*y = *x + *y
	*x = *y - *x
	defer wg.Done()
}

func main() {
	n := 50 // fibonacci N number
	println(fibonacci(n))
}