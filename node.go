package bst

type Node struct {
	Value int;
	Parent, Left, Right *Node;
}

func NewNode(value int) (*Node, error) {
	return &Node {
		Value: value,
	}, nil
}
