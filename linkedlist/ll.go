// A singly linked list
package linkedlist

type Cell struct {
	data int
	next *Cell
}

type List struct {
	top      *Cell
	sentinel Cell
}

func CreateNew() List {
	list := List{}
	list.sentinel = Cell{data: 0, next: nil}
	list.top = &list.sentinel

	return list
}

func (li *List) GetItems() []int {
	ret := []int{}

	for {
		if li.top.next == nil {
			return ret
		}

		ret = append(ret, li.top.next.data)
		li.top = li.top.next
	}
}

func (li *List) Prepend(val int) {
	top := li.top
	n := Cell{data: val, next: top.next}
	top.next = &n
}

func (li *List) Append(val int) {
	top := li.top
	n := Cell{data: val, next: nil}

	for {
		if top.next == nil {
			top.next = &n
			break
		} else {
			top = top.next
		}
	}
}

func (li *List) GetFirst() int {
	return li.top.next.data
}

func (li *List) GetLast() int {
	top := li.top

	for {
		if top.next == nil {
			return top.data
		} else {
			top = top.next
		}
	}
}

func (li *List) GetAndRemoveFirst() int {
	f := li.top.next
	li.top.next = f.next
	return f.data
}

func (li *List) GetAndRemoveLast() int {
	top := li.top
	prev := top

	for {
		if top.next == nil {
			prev.next = nil
			return top.data
		} else {
			prev = top
			top = top.next
		}
	}
}
