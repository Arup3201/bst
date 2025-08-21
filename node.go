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

func (node *Node)InOrder(p []int) ([]int) {
	if node==nil {
		return p
	}

	p = node.Left.InOrder(p)
	p = append(p, node.Value)
	p = node.Right.InOrder(p)

	return p
}

func (node *Node)PreOrder(p []int) ([]int) {
	if node==nil {
		return p
	}

	p = append(p, node.Value)
	p = node.Left.PreOrder(p)
	p = node.Right.PreOrder(p)

	return p
}


func (node *Node) PostOrder (p []int) ([]int) {
	if node==nil {
		return p
	}

	p = node.Left.PostOrder(p)
	p = node.Right.PostOrder(p)
	p = append(p, node.Value)

	return p
}

