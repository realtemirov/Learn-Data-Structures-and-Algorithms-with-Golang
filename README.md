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

    2. [Chapter 2: Getting Started with Go for Data Structures and Algorithms Technical requirements](#chapter-2-getting-started-with-go-for-data-structures-and-algorithms-technical-requirements)


## Hello World !

This is the traditional "Hello World" program in Go. It is the first program that most people write when they are learning a new programming language. The program prints the string "Hello, World!" to the console.
You can run the program below by copying the code into a file named `hello-world.go` and running `go run hello-world.go` in your terminal.

[Code](./HelloWorld/hello_world.go)

![Hello World](./images/hello_world.png)

## Section 1: Introduction to Data Structures and Algorithms and the Go Language

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
[Code](./Chapter01/List/list.go)
![List result](./images/list.png)

## Tuples
[Code](./Chapter01/Tuples/tuples.go)
![Tuples result](./images/tuples.png)

## Heap
[Code](./Chapter01/Heap/heap.go)
![Heap result](./images/heap.png)

## Adapter

![Adapter example](https://refactoring.guru/images/patterns/content/adapter/adapter-en-2x.png)

### Read more about the Adapter pattern:

[Refactoring.guru](https://refactoring.guru/design-patterns/adapter) \
[GolangByExample.com](https://golangbyexample.com/adapter-design-pattern-go/)

[Code](./Chapter01/Adapter/adapter.go)

![Adapter result](./images/adapter.png)

## Bridge

![Bridge example](https://refactoring.guru/images/patterns/content/bridge/bridge-2x.png)

### Read more about the Bridge pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/bridge) \
[GolangByExample.com](https://golangbyexample.com/bridge-design-pattern-in-go/)

[Code](./Chapter01/Bridge/bridge.go)

![Bridge result](./images/bridge.png)

## Composite

![Composite example](https://refactoring.guru/images/patterns/content/composite/composite.png)

### Read more about the Composite pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/composite) \
[GolangByExample.com](https://golangbyexample.com/composite-design-pattern-golang)

[Code](./Chapter01/Composite/composite.go)

![Composite result](./images/composite.png)

## Decorator

![Decorator example](https://refactoring.guru/images/patterns/content/decorator/decorator.png)

### Read more about the Decorator pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/decorator) \
[GolangByExample.com](https://golangbyexample.com/decorator-pattern-golang/)

[Code](./Chapter01/Decorator/decorator.go)

![Decorator result](./images/decorator.png)

## Facade

![Facade example](https://refactoring.guru/images/patterns/content/facade/facade.png)

### Read more about the Facade pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/facade) \
[GolangByExample.com](https://golangbyexample.com/facade-design-pattern-in-golang/)

[Code](./Chapter01/Facade/facade.go)

![Facade result](./images/facade.png)

## Flyweight

![Flyweight example](https://refactoring.guru/images/patterns/content/flyweight/flyweight.png)

### Read more about the Flyweight pattern:
[Refactoring.guru](https://refactoring.guru/design-patterns/flyweight) \
[GolangByExample.com](https://golangbyexample.com/flyweight-design-pattern-golang/)

[Code](./Chapter01/Flyweight/flyweight.go)

![Flyweight result](./images/flyweight.png)

## Private Class

[Code](./Chapter01/PrivateClass/private_class.go)

![Private Class result](./images/private_class.png)

## Chapter 2: Getting Started with Go for Data Structures and Algorithms Technical requirements

Getting Started with Go for Data Structures and Algorithms, covers Go-specific data
structures, such as arrays, slices, two-dimensional slices, maps, structs, and channels.
Variadic functions, deferred function calls, and panic and recover operations are
introduced. Slicing operations, such as enlarging using append and copy, assigning parts,
appending a slice, and appending part of a slice, are also presented in this chapter.

## Contributing

Contributions are welcome! If you have a solution to an exercise that is different from the one provided, feel free to open a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/exercise-solution`).
3. Commit your changes (`git commit -m 'Add solution for Exercise X.Y'`).
4. Push to the branch (`git push origin feature/exercise-solution`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.