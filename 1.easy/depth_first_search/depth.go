package depth

type Node struct {
	Name     string
	Children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		newChild := &Node{
			Name:     name,
			Children: []*Node{},
		}
		n.Children = append(n.Children, newChild)
	}
	return n
}
