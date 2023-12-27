package main

import "fmt"

// Queue represent a basic queue data structure
type Queue struct {
	items []interface{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	front := q.items[0]
	q.items = q.items[1:]
	return front
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	// Create a new queue
	myQueue := Queue{}

	// Enqueue elements
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	// Dequeue and print elements
	for !myQueue.IsEmpty() {
		fmt.Println(myQueue.Dequeue())
	}
}
