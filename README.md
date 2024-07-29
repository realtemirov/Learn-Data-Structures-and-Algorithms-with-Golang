# Learn Data Structures and Algorithms with Golang
Welcome to the repository for the exercises from the book **[Learn Data Structures and Algorithms with Golang](https://www.amazon.com/Go-Programming-Beginner-Professional-everything/dp/1803243058)** by **[Bhagvan Kommadi](https://in.linkedin.com/in/bhagvan-kommadi-b463a6)**. This repository contains solutions for each chapter of the book.

## Table of Contents

0. [Hello World](#hello-world-)

1. [Section 1: Introduction to Data Structures and Algorithms and the Go Language](#section-1-introduction-to-data-structures-and-algorithms-and-the-go-language)
    1. [Chapter 1: Data Structures and Algorithms](#chapter-1-data-structures-and-algorithms)
        * [List](#list)
        * [Tuples](#tuples)
        * [Heap](#heap)
        * [Adapter](#adapter)
        * [Bridge](#bridge)
        * [Composite](#composite)
        * [Decorator](#decorator)
        * [Facade](#facade)
        * [Flyweight](#flyweight)
        * [Private Class](#private-class)
        * [Proxy](#proxy)
        * [Flow chart](#flow-chart)
        * [Pseudo code](#pseudo-code)
        * [Complexity and performance analysis](#complexity-and-performance-analysis)
        * [Complexity](#complexity)
        * [Linear Complexity](#linear-complexity)
        * [Quadratic Complexity](#quadratic-complexity)
		* [Cubic Complexity](#cubic-complexity)
		* [Logarithmic Complexity](#logarithmic-complexity)
        * [Brute force algorithm](#brute-force-algorithms)
		* [Divede and Conquer Algorithms](#divede-and-conquer-algorithms)
		* [Backtracking algorithms](#backtracking-algorithms)

    2. [Chapter 2: Getting Started with Go for Data Structures and Algorithms Technical requirements](#chapter-2-getting-started-with-go-for-data-structures-and-algorithms-technical-requirements)
        * [Arrays](#arrays)
		* [The Len function](#the-len-function)
		* [Slice function](#slice-function)
		* [Two dimensional slices](#two-dimensional-slices)
		* [Maps](#maps)
		* [Database operations](#databasa-operations)
		  * [GetCustomer](#the-getcustomer-function)
		  * [InsertCustomer](#the-insertcustomer-function)
		  * [UpdateCustomer](#the-updatecustomer-function)
		  * [DeleteCustomer](#the-deletecustomer-function)
		* [CRUD web forms](#crud-web-forms)
		  * [HTML Template](#html-template)
		  * [Create function](#create-function)
		  * [Insert function](#insert-function)
		  * [Alter function](#alter-function)
		  * [Update function](#update-function)
		  * [Delete function](#delete-function)
		  * [View function](#view-function)
		  * [The main function](#the-main-function)
		  * [Header template](#header-template)
		  * [Home template](#home-template)
		  * [Footer template](#footer-template)
		  * [Menu template](#menu-template)
		  * [Create template](#create-template)
		  * [Update template](#update-template)
		  * [View template](#view-template)
2. [Section 2: Basic Data Structures and Algorithms using Go](#section-2-basic-data-structures-and-algorithms-using-go)
	
	3. [Chapter 3: Linear Data Structures](#chapter-3-linear-data-structures)
		* [Lists](#lists)
			* [Linked list](#linked-list)
			* [Doubly linked list](#doubly-linked-list)
		* [Sets](#sets)
		* [Tuples](#tuples-1)


## Hello World !

This is the traditional "Hello World" program in Go. It is the first program that most people write when they are learning a new programming language. The program prints the string "Hello, World!" to the console.
You can run the program below by copying the code into a file named `hello-world.go` and running `go run hello-world.go` in your terminal.

[Code](./HelloWorld/hello_world.go)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

![Hello World](./images/hello_world.png)

# Section 1: Introduction to Data Structures and Algorithms and the Go Language

We will be introducing the abstract data types, definition, and classification of data
structures. Readers will be well-versed with performance analysis of algorithms and
choosing appropriate data structures for structural design patterns after reading this part.

## Chapter 1: Data Structures and Algorithms

Data Structures and Algorithms, focuses on the definition of abstract data types,
classifying data structures into linear, non-linear, homogeneous, heterogeneous, and
dynamic types. Abstract data types, such as container, list, set, map, graph, stack, and
queue, are presented in this chapter. This chapter also covers the performance analysis of
data structures, as well as the correct choice of data structures and structural design
patterns

![Data structures](./images/data-structures.png)

## List

[Code](./Chapter01/01-List/list.go)

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
```

![List result](./images/list.png)

## Tuples

[Code](./Chapter01/02-Tuples/tuples.go)

```go
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
```

![Tuples result](./images/tuples.png)

## Heap

[Code](./Chapter01/03-Heap/heap.go)

```go
package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
```

![Heap result](./images/heap.png)

## Adapter

![Adapter example](https://refactoring.guru/images/patterns/content/adapter/adapter-en-2x.png)

### Read more about the Adapter pattern:

[Refactoring.guru](https://refactoring.guru/design-patterns/adapter) \
[GolangByExample.com](https://golangbyexample.com/adapter-design-pattern-go/)

[Code](./Chapter01/04-Adapter/adapter.go)

```go
package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
```

![Adapter result](./images/adapter.png)

## Bridge

![Bridge example](https://refactoring.guru/images/patterns/content/bridge/bridge-2x.png)

### Read more about the Bridge pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/bridge) \
[GolangByExample.com](https://golangbyexample.com/bridge-design-pattern-in-go/)

[Code](./Chapter01/05-Bridge/bridge.go)

```go
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

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = Contour{x: x, y: y}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
```

![Bridge result](./images/bridge.png)

## Composite

![Composite example](https://refactoring.guru/images/patterns/content/composite/composite.png)

### Read more about the Composite pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/composite) \
[GolangByExample.com](https://golangbyexample.com/composite-design-pattern-golang)

[Code](./Chapter01/06-Composite/composite.go)

```go
package main

import "fmt"

type IComposite interface {
	perform()
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}
func (branch *Branch) addLeaf(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet: " + leaf.name)
}

func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}

	branch.addLeaf(leaf1)
	branch.addLeaf(leaf2)
	branch.addBranch(branch2)

	branch.perform()

}
```

![Composite result](./images/composite.png)

## Decorator

![Decorator example](https://refactoring.guru/images/patterns/content/decorator/decorator.png)

### Read more about the Decorator pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/decorator) \
[GolangByExample.com](https://golangbyexample.com/decorator-pattern-golang/)

[Code](./Chapter01/07-Decorator/decorator.go)

```go
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
```

![Decorator result](./images/decorator.png)

## Facade

![Facade example](https://refactoring.guru/images/patterns/content/facade/facade.png)

### Read more about the Facade pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/facade) \
[GolangByExample.com](https://golangbyexample.com/facade-design-pattern-in-golang/)

[Code](./Chapter01/08-Facade/facade.go)

```go
package main

import "fmt"

type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("Account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("Getting account by id")
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("Delete account by id")
}

type Customer struct {
	id   string
	name string
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("Customer creation")
	customer.name = name
	return customer
}

type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(srcAccountId, destAccountId string, amount float32) *Transaction {
	fmt.Println("Transaction creation")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	customer, account = facade.createCustomerAccount("John Doe", "savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("1", "2", 100)
	fmt.Println(transaction.amount)
}
```

![Facade result](./images/facade.png)

## Flyweight

![Flyweight example](https://refactoring.guru/images/patterns/content/flyweight/flyweight.png)

### Read more about the Flyweight pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/flyweight) \
[GolangByExample.com](https://golangbyexample.com/flyweight-design-pattern-golang/)

[Code](./Chapter01/09-Flyweight/flyweight.go)

```go
package main

import "fmt"

type DataTransferObject interface {
	getId() string
}

type Customer struct {
	id   string
	name string
	ssn  string
}

func (customer Customer) getId() string {
	return customer.id
}

type Employee struct {
	id   string
	name string
}

func (employee Employee) getId() string {
	return employee.id
}

type Manager struct {
	id   string
	name string
	dept string
}

func (manager Manager) getId() string {
	return manager.id
}

type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address Address) getId() string {
	return address.id
}

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]

	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)

		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}

	return dto
}

func main() {
	var factory = DataTransferObjectFactory{pool: make(map[string]DataTransferObject)}
	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer ", customer.getId())

	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee ", employee.getId())

	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager ", manager.getId())

	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address ", address.getId())
}
```

![Flyweight result](./images/flyweight.png)

## Private Class

[Code](./Chapter01/10-PrivateClass/private_class.go)

```go
package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *Account) getId() string {
	return account.details.id
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	var account *Account = &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden:", string(jsonAccount))
	fmt.Println("Account ID:", account.getId())
	fmt.Println("Account Type:", account.getAccountType())
}
```

![Private Class result](./images/private_class.png)

## Proxy

![Proxy exampl](https://refactoring.guru/images/patterns/content/proxy/proxy.png)

### Read more about the Proxy pattern:
[Refaactoring.guru](https://refactoring.guru/design-patterns/proxy) \
[GolangByExample.com](https://golangbyexample.com/proxy-design-pattern-in-golang/)

[Code](./Chapter01/11-Proxy/proxy.go)

```go
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
```

![Proxy result](./images/proxy.png)

## Flow chart
Example of Flow chart

![Flow chart](./images/flow_chart.png)

## Pseudo code
Example of pseudo code
```pseudo-code
maximum(arr) {
    n <- len(arr)
    max <- arr[0]
    for k <-0, n do {
        if arr[k] > max {
            max <-arr[k]
        }
    }
    return max
}
```

## Complexity and performance analysis
The complexity is how the algorithm scales when the number
of input parameters increases.

Performance is a measure of time, space, memory, and other
parameters

## Complexity

[Code](./Chapter01/12-Complexity/complexity.go)

```go
package main

import "fmt"

func main() {
	var (
		m [10]int
		k int
	)

	for k = 0; k < 10; k++ {
		m[k] = k + 200
		fmt.Printf("Element [%d] = %d\n", k, m[k])
	}
}
```

![Result of Complexity](./images/complexity.png)

## BigO notation

*T(n)=O(n)*

Using Big O notation, the constant time *O(1)*, linear time *O(n)*, logarithmic time *O(logn)*, cubic time *O(n^3)* and quadratic time *O(n^2)* complexity are different complexity types for an algorithm.

## Linear complexity

[Code](./Chapter01/13-LinearComplexity/linear_complexity.go)

```go
package main

import "fmt"

func main() {
	var (
		m [10]int
		k int
	)

	for k = 0; k < 10; k++ {
		m[k] = k * 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
```

![Result of Linear Complexity](./images/linear_complexity.png)

## Quadratic complexit

[Code](./Chapter01/14-QuadraticComplexity/quadratic_complexity.go)

```go
package main

import "fmt"

func main() {
	var (
		k, l int
	)

	for k = 1; k <= 10; k++ {
		fmt.Println("Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x int = l * k
			fmt.Println(x)
		}
	}
}
```

![Result of Quadratic Complexity](./images/quadratic_complexity.png)

## Cubic Complexity

[Code](./Chapter01/15-CubicComplexity/cubic_complexity.go)

```go
package main

import "fmt"

func main() {
	var k, l, m int

	var arr [10][10][10]int

	for k = 0; k < 10; k++ {
		for l = 0; l < 10; l++ {
			for m = 0; m < 10; m++ {
				arr[k][l][m] = 1
				fmt.Println("Element value ", k, l, m, arr[k][l][m])
			}
		}
	}
}
```

![Result of Cubic complexity](./images/cubic_complexity.png)

## Logarithmic complexity

[Code](./Chapter01/16-LogarithmicComplexity/logarithmic_complexity.go)

```go
package main

import "fmt"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree) {
	if tree != nil {
		fmt.Println("Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

func main() {
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.LeftNode.insert(7)
	print(tree)
}
```

![Result of Logarithmic Complexity](./images/logarithmic_complexity.png)

## Brute force algorithms

[Code](./Chapter01/17-BruteForceAlgorithms/brute_force_algorithms.go)

```go
package main

import "fmt"

func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}

	return false
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 9}
	var check bool = findElement(arr, 10)
	fmt.Println(check)

	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)
}
```

![Result of Brute force algorihtm](./images/brute_force_algorithms.png)

## Divede and Conquer Algorithms

[Code](./Chapter01/18-DivedeAndConquerAlgorithms/divede_and_conquer_algorithms.go)

```go
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
```

![Result of Divede and Conquer Algorithms](./images/divede_and_conquer_algorithms.png)

## Backtracking algorithms

[Code](./Chapter01/19-BacktrackingAlgorithms/backtracking_algorithms.go)

```go
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
```

![Result of Backtracking algorithms](./images/backtracking_algorithms.png)

## Chapter 2: Getting Started with Go for Data Structures and Algorithms Technical requirements

Getting Started with Go for Data Structures and Algorithms, covers Go-specific data
structures, such as arrays, slices, two-dimensional slices, maps, structs, and channels.
Variadic functions, deferred function calls, and panic and recover operations are
introduced. Slicing operations, such as enlarging using append and copy, assigning parts,
appending a slice, and appending part of a slice, are also presented in this chapter.

## Arrays

[Code](./Chapter02/01-Arrays/arrays.go)

```go
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
```

![Result of Arrays](./images/arrays.png)

## The len function

[Code](./Chapter02/02-Slices/basic_slice.go)

```go
package main

import "fmt"

func main() {
	var slice = []int{1, 3, 5, 6}
	slice = append(slice, 8)
	fmt.Println("Capacity", cap(slice))
	fmt.Println("Length", len(slice))
}
```

![Result of function](./images/slice_len_function.png)

## Slice function

[Code](./Chapter02/02-Slices/slices.go)

```go
package main

import "fmt"

func twiceValue(slice []int) {
	var i int
	var value int
	for i, value = range slice {
		slice[i] = 2 * value
	}
}

func main() {
	var slice = []int{1, 3, 5, 6}
	twiceValue(slice)

	var i int
	for i = 0; i < len(slice); i++ {
		fmt.Println("new slice value ", slice[i])
	}
}
```

![Result of function](./images/slices.png)

## Two dimensional slices

[Code](./Chapter02/03-TwoDimensionalSlices/two_dimensional_array.go)

```go
package main

import "fmt"

func main() {
	var twoDArray [8][8]int
	twoDArray[3][6] = 18
	twoDArray[7][4] = 3
	fmt.Println(twoDArray)
}
```

![Result of function](./images/two_dimensional_array.png)

[Code](./Chapter02/03-TwoDimensionalSlices/two_dimensional_slices.go)

```go
package main

import "fmt"

func main() {
	var (
		rows int
		cols int
	)

	rows = 7
	cols = 9
	var twoDSlices = make([][]int, rows)
	var i int
	for i = range twoDSlices {
		twoDSlices[i] = make([]int, cols)
	}
	fmt.Println(twoDSlices)
}
```

![Result of function](./images/two_dimensional_slices.png)

[Code](./Chapter02/03-TwoDimensionalSlices/append_slice.go)

```go
package main

import "fmt"

func main() {
	var arr = []int{5, 6, 7, 8, 9}
	var slice1 = arr[:3]
	fmt.Println("slice1", slice1)
	var slice2 = arr[1:5]
	fmt.Println("slice2", slice2)
	var slice3 = append(slice2, 12)
	fmt.Println("slice3", slice3)
}
```

![Result of function](./images/append_slice.png)

## Maps

[Code](./Chapter02/04-Maps/maps.go)

```go
package main

import "fmt"

func main() {
	var languages = map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}

	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	var i int
	var value string

	for i, value = range languages {
		println("language ", i, ":", value)
	}

	fmt.Println("product with key 2", products[2])

	delete(products, 1)

	fmt.Println(products)
}
```

![Result of Maps](./images/maps.png)

## Databasa operations

[Code](./Chapter02/05-DatabaseOperations/database_operations.go)

### The GetCustomer function

```go
func GetCustomers() []Customer {
	var database *sql.DB = GetConnection()

	var (
		err  error
		rows *sql.Rows
	)

	rows, err = database.Query("SELECT * FROM customers ORDER BY CustomerId DESC")
	if err != nil {
		panic(err.Error())
	}

	var (
		customer  = Customer{}
		customers = []Customer{}
	)

	for rows.Next() {
		var (
			customerId   int
			customerName string
			ssn          string
		)
		err = rows.Scan(
			&customerId,
			&customerName,
			&ssn,
		)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()

	return customers
}
```

![Result of function](./images/get_customer.png)

### The GetCustomerById function

```go
func GetCustomerById(customerId int) Customer {
	var (
		database *sql.DB = GetConnection()
		err      error
		rows     *sql.Rows
	)

	rows, err = database.Query("SELECT * FROM customers WHERE CustomerId = ?", customerId)
	if err != nil {
		panic(err.Error())
	}

	var customer Customer = Customer{}

	for rows.Next() {
		var (
			customerId   int
			customerName string
			ssn          string
		)

		err = rows.Scan(
			&customerId,
			&customerName,
			&ssn,
		)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
	}

	defer database.Close()
	return customer
}
```

### The InsertCustomer function

```go
func InsertCustomer(customer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		insert   *sql.Stmt
	)

	insert, err = database.Prepare("INSERT INTO customers (CustomerName, SSN) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(customer.CustomerName, customer.SSN)
	
	defer database.Close()
}	
```

![Result of function](./images/insert_customer.png)

### The UpdateCustomer function

```go
func UpdateCustomer(customer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		update   *sql.Stmt
	)

	update, err = database.Prepare("UPDATE customers SET CustomerName = ?, SSN = ? WHERE CustomerId = ?")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)

	defer database.Close()
}
```

![Result of function](./images/update_customer.png)

### The DeleteCustomer function

```go
func DeleteCustomer(csutomer Customer) {
	var (
		database *sql.DB = GetConnection()
		err      error
		delete   *sql.Stmt
	)

	delete, err = database.Prepare("DELETE FROM customers WHERE CustomerId = ?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(csutomer.CustomerId)

	defer database.Close()
}
```

![Result of function](./images/delete_customer.png)

## CRUD web forms

[Code](./Chapter02/06-WebForms/web_forms.go)

### HTML Template

[HTML Form](./Chapter02/06-WebForms/main.html)

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(writer http.ResponseWriter, reader *http.Request) {
	var template_html *template.Template = template.Must(template.ParseFiles("./Chapter02/06-WebForms/main.html"))
	template_html.Execute(writer, nil)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
```

```html
<html>
    <body>
        <p>Welcome to Web Forms</p>
    </body>
</html>
```

![Result of function](./images/web_form.png)

![Result of browser](./images/web_form_browser.png)

### Create function

```go
func Create(writer http.ResponseWriter, request *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}
```

### Insert function

```go
func Insert(writer http.ResponseWriter, request *http.Request) {
	var (
		customer  Customer
		customers []Customer
	)

	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	InsertCustomer(customer)

	customers = GetCustomers()

	template_html.ExecuteTemplate(writer, "Home", customers)
}
```

### Alter function

```go
func Alter(writer http.ResponseWriter, request *http.Request) {
	var (
		customer      Customer
		customerId    int
		customerIdStr string
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", customerId)

	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	UpdateCustomer(customer)

	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}
```

### Update function

```go
func Update(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer = GetCustomerById(customerId)

	template_html.ExecuteTemplate(writer, "Update", customer)
}
```

### Delete function

```go
func Delete(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)

	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)

	customers = GetCustomers()

	template_html.ExecuteTemplate(writer, "Home", customers)
}
```

### View function

```go
func View(writer http.ResponseWriter, request *http.Request) {
	var (
		customerId    int
		customerIdStr string
		customer      Customer
		customers     []Customer
	)

	customerIdStr = request.FormValue("id")

	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer = GetCustomerById(customerId)
	fmt.Println(customer)

	customers = GetCustomers()
	customers = append(customers, customer)

	template_html.ExecuteTemplate(writer, "View", customers)
}
```

### The main function

```go
func main() {
	log.Println("Server started on: http://localhost:8000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":8000", nil)
}
```

### Header template

```html
{{ define "Header" }}
<!DOCTYPLE html>
<html>
    <head>
        <title>CRM</title>
        <meta charset="UTF-8" />
    </head>
    <body>
        <h1>Customer Management - CRM</h1>
{{ end }}
```

### Home template

```html
{{ define "Home" }}
    {{ template "Header"}}
    {{ template "Menu"}}
        <h2>Customers</h2>
        <table border="1">
            <thead>
                <tr>
                    <td>CustomerId</td>
                    <td>CustomerName</td>
                    <td>SSN</td>
                    <td>Update</td>
                    <td>Delete</td>
                </tr>
            </thead>
            <tbody>
                {{ range . }}
                    <tr>
                        <td>{{ .CustomerId }}</td>
                        <td>{{ .CustomerName }}</td>
                        <td>{{ .SSN }}</td>
                        <td><a href="/view?id={{ .CustomerId }}">View</a></td>
                        <td><a href="/update?id={{ .CustomerId }}">Update</a></td>
                        <td><a href="/delete?id={{ .CustomerId }}">Delete</a></td>
                    </tr>
                {{ end }}
            </tbody>
        </table>
    {{ template "Footer"}}
{{ end }}
```

### Footer template

```html
{{ define "Footer" }}
<!DICTYPLE html>
    </body>
</html>
{{ end }}
```

### Menu template

```html
{{ define "Menu" }}
    <a href="/">Home</a> | <a href="/create">Create Customer</a>
{{ end }}
```

### Create template

```html
{{ define "Create" }}
    {{ template "Header"}}
    {{ template "Menu"}}
        <br>
            <h1>Create Customer</h1>
        <br>
        <br>
        <form method="POST" action="/insert">
            Customer Name: 
            <input type="text" name="customername" placeholder="Customer name" autofocus/>
            <br>
            <br>
            SSN:
            <input type="text" name="ssn" placeholder="ssn"/>
            <br>
            <br>
            <input type="submit" value="Create Customer" />
        </form>
    {{ template "Footer"}}
{{ end }}
```

### Update template

```html
{{ define "Update" }}
    {{ template "Header"}}
    {{ template "Menu"}}
        <br>
            <h1>Update Customer</h1>
        <br>
        <br>
        <form method="POST" action="/alter">
            <input type="hidden" name="id" value="{{ .CustomerId }}" />
            Customer Name: 
            <input type="text" name="customername" placeholder="Customer name" value="{{ .CustomerName }}" autofocus/>
            <br>
            <br>
            SSN:
            <input type="text" name="ssn" placeholder="ssn" value="{{ .SSN }}"/>
            <br>
            <br>
            <input type="submit" value="Update Customer" />
        </form>
    {{ template "Footer"}}
{{ end }}
```

### View template

```html
{{ define "View" }}
    {{ template "Header"}}
    {{ template "Menu"}}
        <br>
            <h1>View Customer</h1>
        <br>
        <br>
        <table border="1">
            <tr>
                <td>CustomerId</td>
                <td>CustomerName</td>
                <td>SSN</td>
                <td>Update</td>
                <td>Delete</td>
            </tr>
            {{ if . }}
                {{ range . }}
                    <tr>
                        <td>{{ .CustomerId }}</td>
                        <td>{{ .CustomerName }}</td>
                        <td>{{ .SSN }}</td>
                        <td><a href="/update?id={{ .CustomerId }}">Update</a></td>
                        <td><a href="/delete?id={{ .CustomerId }}">Delete</a></td>
                    </tr>
                {{ end }}
            {{ end }}
        </table>
    {{ template "Footer"}}
{{ end }}
```

![Result of function](./images/crm.png)

# Section 2: Basic Data Structures and Algorithms using Go

We will talk about data structures, including linear, non-linear, homogeneous,
heterogeneous, and dynamic types, as well as classic algorithms. This section covers
different types of lists, trees, arrays, sets, dictionaries, tuples, heaps, queues, and stacks,
along with sorting, recursion, searching, and hashing algorithms.

## Chapter 3: Linear Data Structures

Linear Data Structures, covers linear data structures such as lists, sets, tuples,
stacks, and heaps. The operations related to these types, including insertion, deletion,
updating, reversing, and merging are shown with various code samples. In this chapter, we
present the complexity analysis of various data structure operations that display accessing,
search, insertion, and deletion times.

## Lists

### Linked list

[Code](./Chapter03/01-LinkedLists/linked_list.go)

```go
package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node *Node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
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
		node     = &Node{}
		nodeWith *Node
	)
	node.property = property
	node.nextNode = nil
	nodeWith = linkedList.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	var linkedList LinkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}
```

![Result of linked list](./images/linked_list.png)


### Doubly linked list

[Code](./Chapter03/02-DoublyLinkedList/doubly_linked_list.go)

```go
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
```

![Result of doubly linked list](./images/doubly_linked_list.png)

## Sets

[Code](./Chapter03/03-Sets/set.go)

```go
package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(element int) {
	if !set.integerMap[element] {
		set.integerMap[element] = true
	}
}

func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

func main() {
	var set *Set = &Set{}
	set.New()

	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
}
```

![Result of set](./images/set_1.png)

```go
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet *Set = &Set{}
	intersectSet.New()

	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}

	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet *Set = &Set{}
	unionSet.New()

	for value := range set.integerMap {
		unionSet.AddElement(value)
	}

	for value := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

func main() {
	var (
		set        *Set = &Set{}
		anotherSet *Set = &Set{}
	)
	set.New()
	anotherSet.New()

	set.AddElement(1)
	set.AddElement(2)
	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))

}
```

![Result of set](./images/set_2.png)

## Tuples

[Code](./Chapter03/04-Tuples/tuples.go)

```go
package main

import "fmt"

func h(x int, y int) int {
	return x * y
}

func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

func main() {
	fmt.Println(h(g(1, 2)))
}
```

![Result of set](./images/tuples_2.png)

## Contributing

Contributions are welcome! If you have a solution to an exercise that is different from the one provided, feel free to open a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/exercise-solution`).
3. Commit your changes (`git commit -m 'Add solution for Exercise X.Y'`).
4. Push to the branch (`git push origin feature/exercise-solution`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
