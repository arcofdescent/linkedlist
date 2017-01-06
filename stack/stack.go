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
