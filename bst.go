package bst

type BST struct {
	Root *Node;
}

func NewBST() BST {
	return BST{
		Root: nil, 
	}
}
