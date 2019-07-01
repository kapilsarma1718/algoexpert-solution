package ibt

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{
		value: root,
	}

	for _, value := range values {
		tree.Insert(value)
	}

	return tree
}

func (t *BinaryTree) Insert(value int) *BinaryTree {
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		out := queue[0]
		queue = queue[1:]

		if out.left == nil {
			out.left = NewBinaryTree(value)
			break
		} else if out.right == nil {
			out.right = NewBinaryTree(value)
			break
		}

		queue = append(queue, out.left, out.right)

	}

	return t
}

func (t *BinaryTree) InsertAll(values ...int) *BinaryTree {
	for _, value := range values {
		t.Insert(value)
	}
	return t
}

func (t *BinaryTree) InvertedInsert(value int) *BinaryTree {
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		out := queue[0]
		queue = queue[1:]

		if out.right == nil {
			out.right = NewBinaryTree(value)
			break
		} else if out.left == nil {
			out.left = NewBinaryTree(value)
			break
		}

		queue = append(queue, out.right, out.left)

	}

	return t
}

func (t *BinaryTree) InvertedInsertAll(values ...int) *BinaryTree {
	for _, value := range values {
		t.InvertedInsertAll(value)
	}
	return t
}

func (t *BinaryTree) Equals(other *BinaryTree) bool {

	if other == nil || t.value != other.value {
		return false
	}

	if t.left != nil && !t.left.Equals(other.left) {
		return false
	}

	if t.right != nil && !t.right.Equals(other.right) {
		return false
	}

	return true
}

func (t *BinaryTree) Invert() *BinaryTree {
	return t
}
