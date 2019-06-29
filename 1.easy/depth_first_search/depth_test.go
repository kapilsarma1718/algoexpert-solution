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
	tree2 = tt.NewNode("A")
}

func TestDepthFirstSearch(t *testing.T) {
	caseData1 := testCase{
		[]string{},
		[]string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"},
	}

	caseData2 := testCase{
		[]string{},
		[]string{"A"},
	}

	t.Run(fmt.Sprintf("test1()"), func(t *testing.T) {

		result := tree1.DepthFirstSearch(caseData1.arg)

		if !reflect.DeepEqual(result, caseData1.want) {
			t.Errorf("test1() = %v, want : %v", result, caseData1.want)
		}

	})

	t.Run(fmt.Sprintf("test2()"), func(t *testing.T) {

		result := tree2.DepthFirstSearch(caseData2.arg)

		if !reflect.DeepEqual(result, caseData2.want) {
			t.Errorf("test2() = %v, want : %v", result, caseData2.want)
		}

	})
}
