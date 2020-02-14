package main

type item = string

type node struct {
	value item
}

type queue struct {
	items []*node
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) enqueue(n *node) {
	q.items = append(q.items, n)
}

func (q *queue) dequeue() *node {
	lastItemIndex := len(q.items) - 1
	dequeuedItem := q.items[lastItemIndex]
	q.items = q.items[:lastItemIndex]
	return dequeuedItem
}

func (q *queue) front() *node {
	lastItemIndex := len(q.items) - 1
	return q.items[lastItemIndex]
}

func (q *queue) size() int {
	return len(q.items)
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) String() string {
	str := ""
	for _, item := range q.items {
		str += item.value + "\n"
	}
	return str
}
