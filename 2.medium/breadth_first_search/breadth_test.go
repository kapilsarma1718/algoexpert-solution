package bfs_test

import (
	"fmt"
	"reflect"
	"testing"

	tt "algoexpert-solution/2.medium/breadth_first_search"
)

type testCase struct {
	arg  []string
	want []string
}

var tree1 *tt.Node

func init() {
	setupTree1()
}

func setupTree1() {
	tree1 = tt.NewTree("A")
	tree1.AddChildren("B", "C", "D")

	tree1.Children[0].AddChildren("E", "F")
	tree1.Children[0].Children[1].AddChildren("I", "J")

	tree1.Children[2].AddChildren("G", "H")
	tree1.Children[2].Children[0].AddChildren("K")
}

func TestBreadthFirstSearch(t *testing.T) {
	caseData1 := testCase{
		[]string{},
		[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"},
	}

	t.Run(fmt.Sprintf("test1()"), func(t *testing.T) {

		result := tree1.BreadthFirstSearch(caseData1.arg)

		if !reflect.DeepEqual(result, caseData1.want) {
			t.Errorf("test1() = %v, want : %v", result, caseData1.want)
		}

	})

}
