package interview

import "fmt"

// 寻找二叉树两个节点最短距离，手写测试用例
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// 先序初始化二叉树
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

// 先序遍历二叉树
func transNode(root *treeNode) {
	fmt.Printf("%v ", root.Val)
	if root.Left != nil {
		transNode(root.Left)
	}
	if root.Right != nil {
		transNode(root.Right)
	}
}

// 两个节点最短距离
func minDis(root, a, b *treeNode) int {
	root = commonRoot(root, a, b) // 寻找公共祖先
	if root != nil {
		return dis(root, a) + dis(root, b)
	}
	return -1
}

// 求一个节点到祖先的距离
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

// 寻找最近公共祖先
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
