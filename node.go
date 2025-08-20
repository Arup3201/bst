package bst

type Node struct {
	Value int;
	Left, Right *Node;
}

func NewNode(value int) Node {
	return Node {
		Value: value,
	}
}
