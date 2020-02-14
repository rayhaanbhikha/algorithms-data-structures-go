package main

import "fmt"

func main() {
	q := NewQueue()

	q.enqueue(&node{"A"})
	q.enqueue(&node{"B"})

	fmt.Println(q)
	fmt.Println(q.size())
	fmt.Println(q)

	// fmt.Println(q.dequeue().value)
	// fmt.Println(q.dequeue().value)
	// fmt.Println(q.isEmpty())
}
