// A singly linked list
package linkedlist

type Node struct {
	data int
	Next *Node
}

type List []Node

func CreateNew() List {
	list := List{}
	return &list
}

func (l *List) AppendToList(val int) {
	li := Node{data: val, Next: nil}

	if len(l) > 0 {
		l[len(l)-1].Next = &li
	}

	l = append(l, li)
}

/*
func addToBeginningOfList(val int) []Node {
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
