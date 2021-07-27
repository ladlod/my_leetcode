package common_list

func LevelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	tmp := make([]*TreeNode, 0)
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		l := len(tmp)
		r := make([]int, 0)
		for _, node := range tmp {
			r = append(r, node.Val)
			tmp = append(tmp, node.Left)
			tmp = append(tmp, node.Right)
		}
		res = append([][]int{r}, res...)
		tmp = tmp[l:]
	}

	return res
}
