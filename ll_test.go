package linkedlist

import "testing"
import "fmt"

func TestCreate(t *testing.T) {
	CreateNew()
}

func TestAppend(t *testing.T) {
	l := CreateNew()
	l.Append(2)
	l.Append(4)
	l.Append(0)
	fmt.Printf("%+v\n", l)
}

func TestPrepend(t *testing.T) {
	l := CreateNew()
	l.Append(2)
	l.Append(4)
	l.Prepend(0)
	l.Prepend(6)
	fmt.Printf("%+v\n", l)
}
