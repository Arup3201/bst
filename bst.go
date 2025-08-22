package bst

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

func (tree *BST) Insert(z *Node) *Node {
	y := (*Node)(nil)
	x := tree.Root

	for x != nil {
		y = x

		if z.Value > x.Value {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	z.Parent = y
	if y == nil {
		tree.Root = z
	} else if z.Value > y.Value {
		y.Right = z
	} else {
		y.Left = z
	}

	return z
}

func (tree *BST) InsertAll(nodeValues []int) {
	for _, n := range nodeValues {
		node := NewNode(n)
		tree.Insert(node)
	}
}

func (tree *BST) ShiftTree(u, v *Node) {
	if u.Parent == nil {
		tree.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u
	}
}

func (tree *BST) Delete(z *Node) {
	if z.Left == nil {
		tree.ShiftTree(z, z.Right)
	} else if z.Right == nil {
		tree.ShiftTree(z, z.Left)
	} else {
		y, _ := z.Successor() // it is not a leaf node, so it will definitely have a successor
		if y.Parent != z {
			tree.ShiftTree(y, y.Right)
			y.Right = z.Right
			z.Right = y
		}
		tree.ShiftTree(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

func (tree *BST) InOrder(p []int) []int {
	return tree.Root.InOrder(p)
}

func (tree *BST) PreOrder(p []int) []int {
	return tree.Root.PreOrder(p)
}

func (tree *BST) PostOrder(p []int) []int {
	return tree.Root.PostOrder(p)
}

func (tree *BST) Search(t int) (node *Node, ok bool) {
	if tree.Root == nil {
		return (*Node)(nil), false
	}

	x := tree.Root
	for x != nil {
		if x.Value == t {
			return x, true
		} else if t > x.Value {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	return (*Node)(nil), false
}
