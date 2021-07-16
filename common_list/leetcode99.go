package common_list

func recoverTree(root *TreeNode) {
	sortList := proTraverl(root)
	var node1, node2 *TreeNode
	for i := 0; i < len(sortList)-1; i++ {
		if sortList[i] >= sortList[i+1] {
			if node1 == nil {
				node1 = findTreeNode(root, sortList[i])
				node2 = findTreeNode(root, sortList[i+1])
			} else {
				node2 = findTreeNode(root, sortList[i+1])
			}
		}
	}
	if node1 != nil && node2 != nil {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
}

func findTreeNode(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == v {
		return root
	}
	node1 := findTreeNode(root.Left, v)
	if node1 != nil {
		return node1
	}
	node2 := findTreeNode(root.Right, v)
	if node2 != nil {
		return node2
	}
	return nil
}
