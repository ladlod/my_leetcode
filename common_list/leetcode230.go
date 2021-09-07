package common_list

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

// 思路，中序遍历
func kthSmallest(root *TreeNode, k int) int {
	return innerKthSmallest(root, &k)
}

func innerKthSmallest(root *TreeNode, k *int) int {
	if root.Left != nil {
		left := innerKthSmallest(root.Left, k)
		if *k == 0 {
			return left
		}
	}
	*k--
	if *k == 0 {
		return root.Val
	}
	if root.Right != nil {
		right := innerKthSmallest(root.Right, k)
		if *k == 0 {
			return right
		}
	}

	return -1
}
