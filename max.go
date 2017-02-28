package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	top      *Node
	sentinel Node
}

func CreateNew() List {
	list := List{}
	list.sentinel = Node{data: 0, next: nil}
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
	n := Node{data: val, next: top.next}
	top.next = &n
}

func (li *List) Append(val int) {
	top := li.top
	n := Node{data: val, next: nil}

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

type Stack struct {
	list List
}

func CreateNewStack() Stack {
	st := Stack{}
	st.list = CreateNew()
	return st
}

func (st *Stack) Push(val int) {
	st.list.Append(val)
}

func (st *Stack) Pop() int {
	return st.list.GetAndRemoveLast()
}

func (st *Stack) Max() int {
	items := st.list.GetItems()

	max := getMax(items)
	return max
}

func getMax(items []int) int {
	max := items[0]

	for _, val := range items {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {

	s := CreateNewStack()

	var numLines int
	fmt.Scanln(&numLines)

	lines := make([][]string, 0)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < numLines; i++ {
		l, _ := reader.ReadString('\n')
		l = strings.TrimRight(l, "\n")

		ln := []string{l}

		lines = append(lines, ln)
	}

	for _, line := range lines {
		params := strings.Split(strings.Join(line, ""), " ")

		if params[0] == "1" {
			val, _ := strconv.Atoi(params[1])
			s.Push(val)
		} else if params[0] == "2" {
			s.Pop()
		} else if params[0] == "3" {
			fmt.Printf("%d\n", s.Max())
		}
	}

}
