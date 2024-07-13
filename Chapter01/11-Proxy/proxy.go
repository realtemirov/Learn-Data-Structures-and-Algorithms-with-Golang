package main

import "fmt"

type RealObject struct {
}

type IRealObject interface {
	performActin()
}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

type VirtualProxy struct {
	realObject *RealObject
}

func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("VirtualProxy performAction()")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
