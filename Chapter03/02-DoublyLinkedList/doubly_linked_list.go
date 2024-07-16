package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var (
		nodeWith *Node
	)

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty &&
				node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}

	return nodeWith
}

func (linkedList *LinkedList) AddToHead(property int) {
	var (
		node = &Node{}
	)
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var (
		nodeWith *Node
	)

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var (
		node     *Node = &Node{}
		nodeWith *Node
	)
	node.property = property
	node.nextNode = nil

	nodeWith = linkedList.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith

		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var (
		lastNode *Node
	)

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
			break
		}
	}

	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var (
		node     *Node = &Node{}
		lastNode *Node
	)
	node.property = property
	node.nextNode = nil

	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func main() {
	var linkedList LinkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)

	var node *Node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
