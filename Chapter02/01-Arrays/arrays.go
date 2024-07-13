package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println("printing elaments ", arr[i])
	}

	for i, value := range arr {
		fmt.Println("range ", i, value)
	}

	for _, value := range arr {
		fmt.Println("blank range", value)
	}
}
