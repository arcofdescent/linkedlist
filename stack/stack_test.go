package stack

import "testing"

//import "fmt"

func TestCreate(t *testing.T) {
	CreateNew()
}

func TestPushPop(t *testing.T) {
	st := CreateNew()

	st.Push(6)
	st.Push(7)
	st.Push(8)

	if st.Pop() != 8 {
		t.Error(`pop error`)
	}
	if st.Pop() != 7 {
		t.Error(`pop error`)
	}
	if st.Pop() != 6 {
		t.Error(`pop error`)
	}

	st.Push(1)
	if st.Pop() != 1 {
		t.Error(`pop error`)
	}
}
