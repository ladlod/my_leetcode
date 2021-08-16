package common_list

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }
//

// 测试用例格式：

// 简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。

// 邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。

// 给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/clone-graph
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 深度优先
// 广度优先

func cloneGraph(node *PicNode) *PicNode {
	if node == nil {
		return nil
	}

	visitedMap := make(map[*PicNode]*PicNode)
	return deepCloneGraph(visitedMap, node)
}

func deepCloneGraph(visitedMap map[*PicNode]*PicNode, node *PicNode) *PicNode {
	res := &PicNode{
		Val:       node.Val,
		Neighbors: make([]*PicNode, 0),
	}
	visitedMap[node] = res
	for _, nei := range node.Neighbors {
		if old, ok := visitedMap[nei]; !ok {
			res.Neighbors = append(res.Neighbors, deepCloneGraph(visitedMap, nei))
		} else {
			res.Neighbors = append(res.Neighbors, old)
		}
	}
	return res
}

func breathCloneGraph(node *PicNode) *PicNode {
	if node == nil {
		return nil
	}
	visitedMap := make(map[*PicNode]*PicNode)
	queue := []*PicNode{node}
	visitedMap[node] = &PicNode{Val: node.Val, Neighbors: make([]*PicNode, 0)}

	for len(queue) > 0 {
		old := queue[0]
		queue = queue[1:]
		for _, nei := range old.Neighbors {
			if new, ok := visitedMap[nei]; ok {
				visitedMap[old].Neighbors = append(visitedMap[old].Neighbors, new)
			} else {
				queue = append(queue, nei)
				visitedMap[nei] = &PicNode{Val: nei.Val, Neighbors: make([]*PicNode, 0)}
				visitedMap[old].Neighbors = append(visitedMap[old].Neighbors, visitedMap[nei])
			}
		}
	}

	return visitedMap[node]
}
