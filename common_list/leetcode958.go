package common_list

// 给定一个二叉树，确定它是否是一个完全二叉树。

// 百度百科中对完全二叉树的定义如下：

// 若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 层序遍历二叉树，给二叉树结点编号
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return false
}
