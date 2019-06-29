package bfs

type Node struct {
	Name     string
	Children []*Node
}

func NewTree(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := &Node{
			Name:     name,
			Children: []*Node{},
		}
		n.Children = append(n.Children, child)
	}
	return n
}
