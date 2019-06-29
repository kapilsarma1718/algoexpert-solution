package dst

import (
	"encoding/json"
	"fmt"
)

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

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	if len(n.Children) == 0 {
		return array
	}
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}

func (n *Node) PrintTree() {
	res, _ := json.MarshalIndent(n, "", "    ")
	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", string(res))
	fmt.Println("++++++++++++++")
}
