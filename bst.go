package bst

type BST struct {
	Root *Node;
}

func NewBST() (BST, error) {
	return BST{
		Root: nil, 
	}, nil
}
