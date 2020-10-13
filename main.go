package main

import "fmt"

func fib() chan uint64 {
	c := make(chan uint64)
	var i, j uint64
	go func() {
		for i, j = 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()

	return c
}

func main() {
	c := fib()
	for n := 0; n < 50; n++ {
		fmt.Println(<-c)
	}
}
