package common_list

// 给定一个二叉树，确定它是否是一个完全二叉树。

// 百度百科中对完全二叉树的定义如下：

// 若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 中序遍历二叉树给节点编号，左子节点编号为2*i, 右子节点为2*i+1，并记录节点总数，最后大节点编号等于节点总数即为完全二叉树
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	root.Val = 1
	l, max := 1, 1
	midTravlTree(root, &l, &max)
	return l == max
}

func midTravlTree(root *TreeNode, l *int, max *int) {
	if root.Left != nil {
		root.Left.Val = 2 * root.Val
		*l++
		if root.Left.Val > *max {
			*max = root.Left.Val
		}
		midTravlTree(root.Left, l, max)
	}
	if root.Right != nil {
		root.Right.Val = 2*root.Val + 1
		*l++
		if root.Right.Val > *max {
			*max = root.Right.Val
		}
		midTravlTree(root.Right, l, max)
	}
	return
}
