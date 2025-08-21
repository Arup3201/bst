package bst

import (
	"slices"
	"testing"
)

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
}

func TestBSTInorder(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}
	var nodes = []int{20, 15, 10, 9, 25, 21, 11}
	for _, n := range nodes {
		node, _ := NewNode(n)
		tree.Insert(node)
	}

	got := make([]int, 0, 7)
	got = tree.InOrder(got)
	want := []int{9, 10, 11, 15, 20, 21, 25}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}

func TestBSTPreOrder(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}
	var nodes = []int{20, 15, 10, 9, 25, 21, 11}
	for _, n := range nodes {
		node, _ := NewNode(n)
		tree.Insert(node)
	}

	got := make([]int, 0, 7)
	got = tree.PreOrder(got)
	want := []int{20, 15, 10, 9, 11, 25, 21}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}

func TestBSTPostOrder(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}
	var nodes = []int{20, 15, 10, 9, 25, 21, 11}
	for _, n := range nodes {
		node, _ := NewNode(n)
		tree.Insert(node)
	}

	got := make([]int, 0, 7)
	got = tree.PostOrder(got)
	want := []int{9, 11, 10, 15, 21, 25, 20}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}


func TestBSTSearch(t *testing.T) {
	tree, err := NewBST()
	if err != nil {
		t.Errorf("NewBST failed to create root")
		return
	}
	var nodes = []int{20, 15, 10, 9, 25, 21, 11}
	for _, n := range nodes {
		node, _ := NewNode(n)
		tree.Insert(node)
	}

	_, ok := tree.Search(10)
	want := true
	if ok != want {
		t.Errorf("tree.Search()=_, %t, wanted %t", ok, want)
	}
}
