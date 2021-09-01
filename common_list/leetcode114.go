package common_list

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：

// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 将右子树入栈，左子树移到右子树，无子树可移动时出栈，接到右子树
func Flatten(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil {
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left == nil && len(stack) > 0 {
			p.Right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			p.Right = p.Left
			p.Left = nil
		}
		p = p.Right
	}
}

// 思路2：
// O(1)空间复杂度
// 一个节点的右节点的前驱节点为其左子树的最右节点
func FlattenV2(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left != nil {
			q := p.Left
			for ; q.Right != nil; q = q.Right {
			}
			q.Right = p.Right
			p.Right = p.Left
			p.Left = nil
		}
		p = p.Right
	}
}
