package common_list

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 思路：中 右 左 遍历树，记录行号，该行已遍历过则不继续遍历

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	innerRightSideView(root, &res, 0)
	return res
}

func innerRightSideView(root *TreeNode, res *[]int, i int) {
	if len(*res) < i {
		*res = append(*res, root.Val)
	}

	if root.Right != nil {
		innerRightSideView(root.Right, res, i+1)
	}

	if root.Left != nil {
		innerRightSideView(root.Left, res, i+1)
	}

	return
}
