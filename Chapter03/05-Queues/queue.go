package main

// import "fmt"

// type Queue []*Order

// type Order struct {
// 	priority     int
// 	quantity     int
// 	product      string
// 	customerName string
// }

// func (o *Order) New(priority int, quantity int, product string, customerName string) {
// 	o.priority = priority
// 	o.quantity = quantity
// 	o.product = product
// 	o.customerName = customerName
// }

// func (queue *Queue) Add(order *Order) {
// 	if len(*queue) == 0 {
// 		*queue = append(*queue, order)
// 	} else {
// 		var (
// 			appended bool = false
// 		)

// 		for i, addedOrder := range *queue {
// 			if order.priority > addedOrder.priority {
// 				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
// 				appended = true
// 				break
// 			}
// 		}

// 		if !appended {
// 			*queue = append(*queue, order)
// 		}
// 	}
// }

// func main() {
// 	var (
// 		queue         Queue  = make(Queue, 0)
// 		order1        *Order = &Order{}
// 		priority1     int    = 2
// 		quantity1     int    = 20
// 		product1      string = "Computer"
// 		custoemrName1 string = "Greg White"
// 	)

// 	order1.New(priority1, quantity1, product1, custoemrName1)

// 	var (
// 		order2        *Order = &Order{}
// 		priority2     int    = 1
// 		quantity2     int    = 10
// 		product2      string = "Monitor"
// 		custoemrName2 string = "John Smith"
// 	)
// 	order2.New(priority2, quantity2, product2, custoemrName2)

// 	queue.Add(order1)
// 	queue.Add(order2)

// 	for i := 0; i < len(queue); i++ {
// 		fmt.Println(queue[i])
// 	}
// }
