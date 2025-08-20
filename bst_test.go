package bst

import (
	"fmt"
	"testing"
)

func Print(st *Node) {
	if st == nil {
		return
	}

	Print(st.Left)
	fmt.Printf("%d ", st.Value)
	Print(st.Right)
}

func TestBSTInsertRoot(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}

	node, err := NewNode(2)
	if err != nil {
		t.Errorf("NewNode failed to create node")
		return 
	}

	if got:=tree.Insert(node); tree.Root.Value!=got.Value {
		t.Errorf("Insert operation failed to add node at root")
	}
}

func TestBSTInsertInner(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}

	var nodes = []int{20, 15, 10, 9, 25, 21, 11}

	for _, n := range nodes {
		node, err := NewNode(n)
		if err != nil {
			t.Errorf("NewNode failed to create node")
			return 
		}
		tree.Insert(node)
	}

	Print(tree.Root);
}
