package common_list

import (
	"strconv"
	"strings"
)

// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 中序遍历二叉树，遇到空节点记录一个null，将二叉树组成一个字符串数组
// eg:   1
//     2   3    -> 1,2,4,null,null,null,3,5,null,null,6,null,null
//   4    5  6
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	return strings.Join(this.midTravl(root), ",")
}

func (this *Codec) midTravl(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		res = append(res, "null")
		return res
	} else {
		res = append(res, strconv.Itoa(root.Val))
	}
	res = append(res, this.midTravl(root.Left)...)
	res = append(res, this.midTravl(root.Right)...)

	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	if len(list) <= 0 {
		return nil
	}

	res, _ := this.buildTree(list)
	return res
}

func (this *Codec) buildTree(vals []string) (*TreeNode, []string) {
	if len(vals) <= 0 {
		return nil, vals
	}
	if vals[0] == "null" {
		vals = vals[1:]
		return nil, vals
	}
	val, _ := strconv.Atoi(vals[0])
	res := &TreeNode{
		Val: val,
	}
	vals = vals[1:]
	res.Left, vals = this.buildTree(vals)
	res.Right, vals = this.buildTree(vals)

	return res, vals
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
