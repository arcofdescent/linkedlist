// A queue
package queue

import "github.com/arcofdescent/linkedlist/linkedlist"

type Queue struct {
	list linkedlist.List
}

func CreateNew() Queue {
	q := Queue{}
	q.list = linkedlist.CreateNew()
	return q
}

func (q *Queue) Enqueue(val int) {
	q.list.Append(val)
}

func (q *Queue) Dequeue() int {
	return q.list.GetAndRemoveFirst()
}
