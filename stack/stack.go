// A stack
package stack

import "github.com/arcofdescent/linkedlist/linkedlist"

type Stack struct {
	list linkedlist.List
}

func CreateNew() Stack {
	st := Stack{}
	st.list = linkedlist.CreateNew()
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
