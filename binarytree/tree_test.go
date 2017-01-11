package tree

import "testing"
import "fmt"

func TestCreate(t *testing.T) {
	root := CreateRoot(5)

	if root.Data != 5 {
		t.Error("CreateRoot() error")
	}
}

func TestAddNode(t *testing.T) {
	root := CreateRoot(5)
	root.AddNode(4)
}

func TestPrint(t *testing.T) {
	root := CreateRoot(5)
	root.AddNode(4)
	root.AddNode(6)
	root.AddNode(2)
	root.AddNode(8)
	root.AddNode(1)
	root.AddNode(13)
	root.PrintTree()
	fmt.Println()
}
