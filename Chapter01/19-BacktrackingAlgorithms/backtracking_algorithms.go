package main

import "fmt"

func findElementWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0

	if addValue > k {
		return -1
	}

	if addValue == k {
		num++
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d, ", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		findElementWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l++
	}
	return num
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum = 18
	var combinations [19]int
	findElementWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}
