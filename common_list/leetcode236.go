package common_list

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/*
思路：如果p，q为root本身，则root就是最近公共祖先
     否则，如果p，q均是root的子节点，且不在同一子树，则root为最近公共祖先
	 递归
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
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
