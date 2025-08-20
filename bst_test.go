package bst

import "testing"

func TestNewBST(t *testing.T) {
	_, err := NewBST()
	if err != nil {
		t.Errorf("BST is nil, expected structure of `BST` struct")
	}
}
