package main

import "fmt"

type IProcess interface {
	process()
}

type Process struct{}

func (process Process) process() {
	fmt.Println("Process process")
}

type ProcessDecorator struct {
	processInstance *Process
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	var process = &Process{}
	var decorator = &ProcessDecorator{}

	decorator.process()
	decorator.processInstance = process
	decorator.process()

}
