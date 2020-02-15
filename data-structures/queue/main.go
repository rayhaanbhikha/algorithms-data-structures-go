package queue

import "github.com/rayhaanbhikha/algorithms-data-structures-go/data-structures/node"

// Queue ...
type Queue struct {
	items []*node.SNode
}

// Enqueue ... add item to the queue
func (q *Queue) Enqueue(n *node.SNode) {
	q.items = append(q.items, n)
}

// Dequeue ... pop item off the queue
func (q *Queue) Dequeue() *node.SNode {
	deQueuedItem := q.items[0]
	q.items = q.items[1:]
	return deQueuedItem
}

// Front ... displays first item on the queue
func (q *Queue) Front() *node.SNode {
	lastItemIndex := len(q.items) - 1
	return q.items[lastItemIndex]
}

// Size ...
func (q *Queue) Size() int {
	return len(q.items)
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) String() string {
	str := ""
	for _, item := range q.items {
		str += item.Value + "\n"
	}
	return str
}
