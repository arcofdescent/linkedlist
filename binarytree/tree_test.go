package tree

import "testing"
import "fmt"
import "math/rand"

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

func BenchmarkAddNode(b *testing.B) {
	root := CreateRoot(2)

	for i := 0; i < b.N; i++ {
		d := rand.Intn(1000)
		//fmt.Printf("%d\n", d)
		root.AddNode(d)
	}

	//root.PrintTree()
	//fmt.Println()
}
