package common_list

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	for i, v := range inorder {
		if v == preorder[0] {
			root.Left = BuildTree(preorder[1:i+1], inorder[:i])
			root.Right = BuildTree(preorder[i+1:], inorder[i+1:])
		}
	}

	return root
}
