package main

import (
	"errors"
)

type Queue struct {
	items []interface{}
}

//enqueue
func (q *Queue) Enqueue(v interface{}) {
	q.items = append(q.items, v)
}

// dequeue
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return -1, errors.New("empty queue")
	}

	p := q.items[0]
	q.items = q.items[1:len(q.items)]
	return p, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) <= 0
}

// func main() {
// 	queue := &Queue{}

// 	queue.Enqueue(6)
// 	queue.Enqueue(743)
// 	queue.Enqueue(643)
// 	queue.Enqueue(32)
// 	queue.Enqueue(7)
// 	fmt.Println(queue.items)

// 	queue.Dequeue()
// 	queue.Dequeue()
// 	fmt.Println(queue.items)
// }
