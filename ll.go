// A singly linked list
package linkedlist

//import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	items []Node
}

func CreateNew() List {
	list := List{}
	return list
}

func (l *List) Append(val int) {
	li := Node{data: val, next: nil}

	if len(l.items) > 0 {
		l.items[len(l.items)-1].next = &li
	}

	l.items = append(l.items, li)
}

func (l *List) Prepend(val int) {
	li := Node{data: val, next: nil}

	if len(l.items) == 0 {
		l.items = append(l.items, li)
	} else {
		li_c := []Node{}

		li_c = append(li_c, li)

		for _, n := range l.items {
			li_c = append(li_c, n)
		}

		l.items = li_c
		l.items[0].next = &l.items[1]
	}
}

/*
func getFirst() int {
	f_val := myList[0].data
	return f_val
}

func getLast() int {
	f_val := myList[len(myList)-1].data
	return f_val
}

func getAndRemoveFirst() int {
	f_val := myList[0].data
	myList = myList[1:]
	return f_val
}

func getAndRemoveLast() int {
	f_val := myList[len(myList)-1].data
	myList = myList[0 : len(myList)-1]
	return f_val
}
*/
