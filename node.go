package bst

type Node struct {
	Value               int
	Parent, Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (node *Node) InOrder(p []int) []int {
	if node == nil {
		return p
	}

	p = node.Left.InOrder(p)
	p = append(p, node.Value)
	p = node.Right.InOrder(p)

	return p
}

func (node *Node) PreOrder(p []int) []int {
	if node == nil {
		return p
	}

	p = append(p, node.Value)
	p = node.Left.PreOrder(p)
	p = node.Right.PreOrder(p)

	return p
}

func (node *Node) PostOrder(p []int) []int {
	if node == nil {
		return p
	}

	p = node.Left.PostOrder(p)
	p = node.Right.PostOrder(p)
	p = append(p, node.Value)

	return p
}

func (node *Node) Minimum() *Node { /* Finds the minimum of the subtree with `node` as it's root */
	x := node
	for x.Left != nil {
		x = x.Left
	}

	return x
}

func (node *Node) Maximum() *Node { /* Finds the maximum of the subtree with `node` as it's root */
	x := node
	for x.Right != nil {
		x = x.Right
	}

	return x
}

func (node *Node) Predecessor() (*Node, bool) {
	if node.Left != nil {
		return node.Left.Maximum(), true
	}

	x := node
	y := x.Parent
	for y != nil && y.Left == node {
		x = y
		y = x.Parent
	}

	if y != nil {
		return y, true
	} else {
		return (*Node)(nil), false
	}
}

func (node *Node) Successor() (*Node, bool) {
	if node.Right != nil {
		return node.Right.Minimum(), true
	}

	x := node
	y := x.Parent
	for y != nil && y.Right == node {
		x = y
		y = x.Parent
	}

	if y != nil {
		return y, true
	} else {
		return (*Node)(nil), false
	}
}
