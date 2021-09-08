package common_list

// 返回与给定的前序和后序遍历匹配的任何二叉树。

//  pre 和 post 遍历中的值是不同的正整数。

// 示例：

// 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 首先，我们可以确认的是，根据前序遍历和后序遍历无法确认一棵唯一的二叉树，如 [1,2] [2,1] 这两序列，仅可以确认一棵根节点为1，叶子结点为2的树
// 所以，我们默认所有子树优先放在二叉树的左子树
// 根据先序队列，可以得出先序的第1位为二叉树的根节点，第2位为左子树的根节点，根据后序队列可以找出左子树根节点的位置，其后的串即为右子树
// 以此思路可以递归构造二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) > 1 {
		prel := 1
		postl := 0
		for ; postl < len(postorder); postl++ {
			if preorder[prel] == postorder[postl] {
				break
			}
		}
		root.Left = constructFromPrePost(preorder[prel:prel+postl+1], postorder[:postl+1])
		root.Right = constructFromPrePost(preorder[prel+postl+1:], postorder[postl+1:len(postorder)-1])
	}

	return root
}
