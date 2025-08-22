# Binary Search Tree

BST structure:

```go
type BST struct {
	Root *Node;
}
```

BST Node structure:
```go
type Node struct {
	Value int;
	Parent, Left, Right *Node;
}
```
