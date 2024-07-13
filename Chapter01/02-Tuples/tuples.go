package main

import (
	"fmt"
)

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeriesNamed(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a
	return
}

func powerSeriesError(a int) (int, int, error) {
	var square int = a * a
	var cube int = a * a * a
	return square, cube, nil
}

func main() {
	var square int
	var cube int

	square, cube = powerSeries(3)

	fmt.Println("Square", square, "Cube", cube)
	fmt.Println(powerSeriesNamed(4))
	fmt.Println(powerSeriesError(5))
}
