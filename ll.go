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

func (l *List) AppendTo(val int) {
	li := Node{data: val, next: nil}

	if len(l.items) > 0 {
		l.items[len(l.items)-1].next = &li
	}

	l.items = append(l.items, li)
}

/*
func (l *List) PrependTo(val int) {
	li := Node{data: val, next: &myList[0]}

	var list []Node
	list = append(list, li)
	list = append(list, myList[:]...)

	return list
}

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
