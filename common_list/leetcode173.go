package common_list

// 实现一个二叉搜索树迭代器类BSTIteratorV1 ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
// BSTIteratorV1(TreeNode root) 初始化 BSTIteratorV1 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
// boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
// int next()将指针向右移动，然后返回指针处的数字。
// 注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。

// 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路1：将树中序遍历写成数组进行扁平化处理
type BSTIteratorV1 struct {
	tree  []int
	index int
}

func ConstructorBSTV1(root *TreeNode) BSTIteratorV1 {
	return BSTIteratorV1{
		tree:  transBSTToArray(root),
		index: -1,
	}
}

func (this *BSTIteratorV1) Next() int {
	this.index++
	if this.index >= 0 && this.index < len(this.tree) {
		return this.tree[this.index]
	} else {
		return 0
	}
}

func (this *BSTIteratorV1) HasNext() bool {
	return this.index+1 < len(this.tree)
}

func transBSTToArray(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, transBSTToArray(root.Left)...)
	res = append(res, root.Val)
	res = append(res, transBSTToArray(root.Right)...)
	return res
}

// 思路2：用栈存储节点, next时中序遍历树
// if 栈顶元素左子节点非nil 栈顶元素左子节点入栈，并将该节点左子节点置空
// if 栈顶元素左子节点为nil，出栈该节点，输出值，若该元素右子节点非nil，入栈其右子节点
type BSTIteratorV2 struct {
	root  *TreeNode
	stack []*TreeNode
}

func ConstructorBSTV2(root *TreeNode) BSTIteratorV2 {
	return BSTIteratorV2{
		root:  root,
		stack: []*TreeNode{root},
	}
}

func (this *BSTIteratorV2) Next() int {
	for len(this.stack) > 0 {
		if this.stack[len(this.stack)-1].Left != nil {
			this.stack = append(this.stack, this.stack[len(this.stack)-1].Left)
			this.stack[len(this.stack)-2].Left = nil
		} else {
			tmp := this.stack[len(this.stack)-1]
			if tmp.Right != nil {
				this.stack[len(this.stack)-1] = tmp.Right
			} else {
				this.stack = this.stack[:len(this.stack)-1]
			}
			return tmp.Val
		}
	}

	return 0
}

func (this *BSTIteratorV2) HasNext() bool {
	return len(this.stack) > 0
}
