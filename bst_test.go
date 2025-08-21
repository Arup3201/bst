package bst

import (
	"slices"
	"testing"
)

func TestBSTInorder(t *testing.T) {
	tree := NewBST()
	tree.InsertAll([]int{20, 15, 10, 9, 25, 21, 11})

	got := make([]int, 0, 7)
	got = tree.InOrder(got)
	want := []int{9, 10, 11, 15, 20, 21, 25}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}

func TestBSTPreOrder(t *testing.T) {
	tree := NewBST()
	tree.InsertAll([]int{20, 15, 10, 9, 25, 21, 11})


	got := make([]int, 0, 7)
	got = tree.PreOrder(got)
	want := []int{20, 15, 10, 9, 11, 25, 21}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}

func TestBSTPostOrder(t *testing.T) {
	tree := NewBST()
	tree.InsertAll([]int{20, 15, 10, 9, 25, 21, 11})


	got := make([]int, 0, 7)
	got = tree.PostOrder(got)
	want := []int{9, 11, 10, 15, 21, 25, 20}
	if !slices.Equal(got, want) {
		t.Errorf("tree.Inorder()=%v, expected %v", got, want)
	}
}


func TestBSTSearch(t *testing.T) {
	tree := NewBST()
	tree.InsertAll([]int{20, 15, 10, 9, 25, 21, 11})

	_, ok := tree.Search(10)
	want := true
	if ok != want {
		t.Errorf("tree.Search()=_, %t, wanted %t", ok, want)
	}
}
