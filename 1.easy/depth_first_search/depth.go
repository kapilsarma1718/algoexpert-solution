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
