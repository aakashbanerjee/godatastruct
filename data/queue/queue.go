package queue

import (
	"fmt"
	"sync"
)

//Queue data struct for FIFO queue
type Queue struct {
	lock sync.Mutex
	data []interface{}
	head int
}

//Enqueue adds data to the queue
func (q *Queue) Enqueue(d interface{}) {
	q.lock.Lock()
	q.data = append(q.data, d)
	q.lock.Unlock()
}

//Display returns slice that holds the queue data and the head
func (q *Queue) Display() ([]interface{}, int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.data, q.head
}

//Dequeue returns the first element(FIFO) and queue empty status
func (q *Queue) Dequeue() (interface{}, bool) {
	if !q.IsEmpty() {
		q.lock.Lock()
		fmt.Printf("head: %v\n", q.head)
		r := q.data[q.head]
		q.data = q.data[q.head+1:]
		defer q.lock.Unlock()
		return r, false
	}
	return nil, true
}

//IsEmpty checks whether queue is empty
func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	if len(q.data) <= 0 {
		defer q.lock.Unlock()
		return true
	}
	defer q.lock.Unlock()
	return false
}
