package common_list

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
