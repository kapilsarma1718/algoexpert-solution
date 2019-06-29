package depth

type Node struct {
	name     string
	children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		name:     name,
		children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		newChild := &Node{
			name:     name,
			children: []*Node{},
		}
		n.children = append(n.children, newChild)
	}
	return n
}
