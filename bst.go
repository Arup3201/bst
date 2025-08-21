package bst

type BST struct {
	Root *Node;
}

func NewBST() (*BST, error) {
	return &BST{
		Root: nil, 
	}, nil
}

func (tree *BST) Insert(z *Node) (*Node){
	y := (*Node)(nil)
	x := tree.Root

	for x != nil {
		y = x

		if z.Value>x.Value {
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

func (tree *BST) InOrder(p []int) ([]int) {
	return tree.Root.InOrder(p)
}

func (tree *BST) PreOrder(p []int) ([]int) {
	return tree.Root.PreOrder(p)
}

func (tree *BST) PostOrder(p []int) ([]int) {
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

