package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct{}

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type Contour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (contour Contour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}
func (contour Contour) resizeByFactor(factor int) {
	contour.factor = factor
}

// type ellipse struct {
// 	a, b, r int
// }

// func (e ellipse) drawShape() {

// }

// func (e ellipse) drawContour() {

// }

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = Contour{x: x, y: y}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
