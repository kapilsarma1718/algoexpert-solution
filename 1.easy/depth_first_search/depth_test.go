package dst_test

import (
	"fmt"
	"reflect"

	"testing"

	tt "algoexpert-solution/1.easy/depth_first_search"
)

type testCase struct {
	arg  []string
	want []string
}

var tree1 *tt.Node
var tree2 *tt.Node

func init() {
	setupTree1()
	setupTree2()
}

func setupTree1() {
	tree1 = tt.NewNode("A")
	tree1.AddChildren("B", "C", "D")

	tree1.Children[0].AddChildren("E", "F")
	tree1.Children[0].Children[1].AddChildren("I", "J")

	tree1.Children[2].AddChildren("G", "H")
	tree1.Children[2].Children[0].AddChildren("K")
}

func setupTree2() {
	tree2 = tt.NewNode("")
}

func TestDepthFirstSearch(t *testing.T) {
	caseData := testCase{
		[]string{},
		[]string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"},
	}

	t.Run(fmt.Sprintf("test1()"), func(t *testing.T) {

		result := tree1.DepthFirstSearch(caseData.arg)

		if !reflect.DeepEqual(result, caseData.want) {
			t.Errorf("test1() = %v, want : %v", result, caseData.want)
		}

	})
}
