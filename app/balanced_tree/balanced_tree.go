package balanced_tree

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return node.Height
}

func balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return height(node.Left) - height(node.Right)
}

func rotateLeft(node *TreeNode, n int) (*TreeNode, int) {
	right := node.Right
	node.Right = right.Left
	right.Left = node

	node.Height = max(height(node.Right), height(node.Left)) + 1
	right.Height = max(height(right.Right), height(right.Left)) + 1

	return right, n + 3
}

func rotateRight(node *TreeNode, n int) (*TreeNode, int) {
	left := node.Left
	node.Left = left.Right
	left.Right = node

	node.Height = max(height(node.Right), height(node.Left)) + 1
	left.Height = max(height(left.Right), height(left.Left)) + 1

	return left, n + 3
}

func rebalance(node *TreeNode, val int, n int) (*TreeNode, int) {
	balance := balanceFactor(node)

	if balance > 1 && val < node.Left.Val {
		return rotateRight(node, n)
	}

	if balance < -1 && val > node.Right.Val {
		return rotateLeft(node, n)
	}

	if balance > 1 && val > node.Left.Val {
		t := 0
		node.Left, t = rotateLeft(node.Left, n)

		return rotateRight(node, n+t)
	}

	if balance < -1 && val < node.Right.Val {
		t := 0
		node.Right, t = rotateRight(node.Right, n)

		return rotateLeft(node, n+t)
	}

	return node, n
}

func minVal(node *TreeNode) int {
	mv := node.Val

	for node.Left != nil {
		mv = node.Left.Val
		node = node.Left
	}

	return mv
}

func Add(node *TreeNode, val int, n int) (*TreeNode, int) {
	if node == nil {
		return &TreeNode{val, nil, nil, 1}, 1
	}

	if node.Val > val {
		node.Left, n = Add(node.Left, val, n)
	} else if node.Val < val {
		node.Right, n = Add(node.Right, val, n)
	} else {
		return node, n
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	return rebalance(node, val, n)
}

func Build(values []int) (*TreeNode, int) {
	var root *TreeNode

	n, t := 0, 0
	for _, val := range values {
		root, n = Add(root, val, 0)
		t += n
	}

	return root, t
}

func Find(root *TreeNode, val int, n int) (*TreeNode, int) {
	if root == nil {
		return nil, n
	}

	if root.Val == val {
		return root, n + 1
	}

	if root.Val > val {
		return Find(root.Left, val, n+1)
	}

	if root.Val < val {
		return Find(root.Right, val, n+1)
	}

	return nil, n + 1
}

func Delete(root *TreeNode, val int, n int) (*TreeNode, int) {
	if root == nil {
		return nil, n
	}

	if root.Val > val {
		return Delete(root.Left, val, n+1)
	}

	if root.Val < val {
		return Delete(root.Right, val, n+1)
	}

	if root.Val == val {
		if root.Left == nil {
			return root.Right, n
		}

		if root.Right == nil {
			return root.Left, n
		}

		root.Val = minVal(root)
		root.Right, n = Delete(root.Right, root.Val, n+1)
	}

	return root, n
}
