package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	Value int
}
type Stack struct {
	elements     []*Element
	elementCount int
}

func (element *Element) String() string {
	return strconv.Itoa(element.Value)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount++
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var (
		length  int      = len(stack.elements)
		element *Element = stack.elements[length-1]
	)

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}

	stack.elementCount = len(stack.elements)

	return element
}

func main() {
	var (
		stack    *Stack   = &Stack{}
		element1 *Element = &Element{3}
		element2 *Element = &Element{5}
		element3 *Element = &Element{7}
		element4 *Element = &Element{9}
	)
	stack.New()

	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
