package common_list

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/validate-binary-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：前序遍历二叉树，判断结果是否有序

func isValidBST(root *TreeNode) bool {
	sortList := proTraverl(root)
	for i := 0; i < len(sortList)-1; i++ {
		if sortList[i] >= sortList[i+1] {
			return false
		}
	}
	return true
}

func proTraverl(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, proTraverl(root.Left)...)
	res = append(res, root.Val)
	res = append(res, proTraverl(root.Right)...)
	return res
}
