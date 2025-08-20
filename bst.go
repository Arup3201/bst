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
