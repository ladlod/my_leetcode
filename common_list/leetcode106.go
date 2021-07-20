package common_list

func BuildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			root.Left = BuildTreePost(inorder[:i], postorder[:i])
			root.Right = BuildTreePost(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}

	return root
}
