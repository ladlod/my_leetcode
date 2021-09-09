package interview

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func initNode(v *[]int) *treeNode {
	if len(*v) <= 0 {
		return nil
	}
	if (*v)[0] == -1 {
		*v = (*v)[1:]
		return nil
	}
	node := &treeNode{
		Val: (*v)[0],
	}
	*v = (*v)[1:]
	node.Left = initNode(v)
	node.Right = initNode(v)
	return node
}

func transNode(root *treeNode) {
	fmt.Printf("%v ", root.Val)
	if root.Left != nil {
		transNode(root.Left)
	}
	if root.Right != nil {
		transNode(root.Right)
	}
}

func minDis(root, a, b *treeNode) int {
	root = commonRoot(root, a, b) // 寻找公共祖先
	if root != nil {
		return dis(root, a) + dis(root, b)
	}
	return -1
}

func dis(root, a *treeNode) int {
	if root == a {
		return 0
	}
	if root.Left != nil {
		l := dis(root.Left, a)
		if l != -1 {
			return l + 1
		}
	}
	if root.Right != nil {
		r := dis(root.Right, a)
		if r != -1 {
			return r + 1
		}
	}
	return -1
}

func commonRoot(root, a, b *treeNode) *treeNode {
	if root == nil {
		return nil
	}
	if root == a || root == b {
		return root
	}

	l := commonRoot(root.Left, a, b)
	r := commonRoot(root.Right, a, b)
	if l != nil && r != nil {
		return root
	} else if l == nil {
		return r
	} else if r == nil {
		return l
	} else {
		return nil
	}
}
