package main

import "fmt"

func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}

	return fibonacci(k-1) + fibonacci(k-2)
}

func main() {
	var i int = 5
	for i = 0; i < 8; i++ {
		var fib = fibonacci(i)
		fmt.Println(fib)
	}
}
