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

func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}

	for len(queue) > 0 {
		out := queue[0]
		queue = queue[1:]

		array = append(array, out.Name)

		for _, child := range out.Children {
			queue = append(queue, child)
		}
	}

	return array
}
