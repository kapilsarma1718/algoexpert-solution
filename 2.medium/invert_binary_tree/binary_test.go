package ibt_test

import (
	"fmt"
	"testing"

	tt "algoexpert-solution/2.medium/invert_binary_tree"
)

func TestTree1(t *testing.T) {
	t.Run(fmt.Sprintf("InvertBinaryTree()"), func(t *testing.T) {

		tree1 := tt.NewBinaryTree(1).InsertAll(2, 3, 4)
		tree1.InvertBinaryTree()

		tree2 := tt.NewBinaryTree(2).InvertedInsertAll(2, 3, 4)

		if !tree1.Equals(tree2) {

			t.Errorf("(InvertBinaryTree) = %v, want : %v", tree1, tree2)
		}

	})
}
