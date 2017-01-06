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

func (l *List) GetItems() []int {
	ret := []int{}

	for _, n := range l.items {
		ret = append(ret, n.data)
	}

	return ret
}

func (l *List) GetFirst() int {
	return l.items[0].data
}

func (l *List) GetLast() int {
	return l.items[len(l.items)-1].data
}

func (l *List) GetAndRemoveFirst() int {
	f_val := l.items[0].data
	l.items = l.items[1:]
	return f_val
}

func (l *List) GetAndRemoveLast() int {
	f_val := l.items[len(l.items)-1].data
	l.items = l.items[0 : len(l.items)-1]
	return f_val
}
