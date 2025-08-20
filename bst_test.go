package bst

import "testing"

func TestBSTInsert(t *testing.T) {
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
