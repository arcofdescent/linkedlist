// sorted binary tree
package tree

import "fmt"

type Node struct {
	Data       int
	LeftChild  *Node
	RightChild *Node
}

func CreateRoot(val int) *Node {
	root := Node{}
	root.Data = val
	return &root
}

func (n *Node) AddNode(val int) {
	if val < n.Data {
		if n.LeftChild == nil {
			n.LeftChild = &Node{Data: val}
		} else {
			n.LeftChild.AddNode(val)
		}
	} else {
		if n.RightChild == nil {
			n.RightChild = &Node{Data: val}
		} else {
			n.RightChild.AddNode(val)
		}
	}
}

func (n *Node) PrintTree() {

	// inorder traversal

	// Process LeftChild
	if n.LeftChild != nil {
		n.LeftChild.PrintTree()
	}

	// Process Node
	fmt.Printf("%d ", n.Data)

	// Process RightChild
	if n.RightChild != nil {
		n.RightChild.PrintTree()
	}

}
