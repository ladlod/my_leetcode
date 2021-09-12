package week_contest

// 有一棵根节点为 0 的 家族树 ，总共包含 n 个节点，节点编号为 0 到 n - 1 。给你一个下标从 0 开始的整数数组 parents ，其中 parents[i] 是节点 i 的父节点。由于节点 0 是 根 ，所以 parents[0] == -1 。

// 总共有 105 个基因值，每个基因值都用 闭区间 [1, 105] 中的一个整数表示。给你一个下标从 0 开始的整数数组 nums ，其中 nums[i] 是节点 i 的基因值，且基因值 互不相同 。

// 请你返回一个数组 ans ，长度为 n ，其中 ans[i] 是以节点 i 为根的子树内 缺失 的 最小 基因值。

// 节点 x 为根的 子树 包含节点 x 和它所有的 后代 节点。

// 输入：parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
// 输出：[7,1,1,4,2,1]
// 解释：每个子树答案计算结果如下：
// - 0：子树内包含节点 [0,1,2,3,4,5] ，基因值分别为 [5,4,6,2,1,3] 。7 是缺失的最小基因值。
// - 1：子树内包含节点 [1,2] ，基因值分别为 [4,6] 。 1 是缺失的最小基因值。
// - 2：子树内只包含节点 2 ，基因值为 6 。1 是缺失的最小基因值。
// - 3：子树内包含节点 [3,4,5] ，基因值分别为 [2,1,3] 。4 是缺失的最小基因值。
// - 4：子树内只包含节点 4 ，基因值为 1 。2 是缺失的最小基因值。
// - 5：子树内只包含节点 5 ，基因值为 3 。1 是缺失的最小基因值。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/smallest-missing-genetic-value-in-each-subtree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 暴力解法，遍历所有节点，记录所有节点及其子树出现过的基因，对每个节点的序列排序并求最小未出现的基因，时间复杂路O(n^2*logn)，超时
// 出现在一个节点的值，仅会出现在它及它到根节点的路径上，记录每条路径上出现过的最小数字即可
func smallestMissingValueSubtreeV1(parents []int, nums []int) []int {
    flag := make([][]int, len(parents))
    for i := 0; i < len(flag); i++ {
        flag[i] = make([]int, 0)
    }
    
    for i := 0; i < len(parents); i++ {
        for j := i; j != -1; j = parents[j] {
            flag[j] = append(flag[j], nums[i]) // 转成压缩列表？
        }
    }
    
    // fmt.Println(flag)
    
    res := make([]int, len(parents))
    for i := 0; i < len(flag); i++ {
        sort.Ints(flag[i])
        j := 1
        for ; j <= len(flag[i]);j++ {
            if flag[i][j-1] != j {
                break
            }
        }
        res[i] = j
    }
    
    return res
}

// 思路
// 
func smallestMissingValueSubtreeV2(parents []int, nums []int) []int {

}

// 输入：parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
// 输出：[7,1,1,4,2,1]
// 找到基因中为1的节点，为找到则结果为[1]
// 从基因为1的节点反向深度优先遍历，遍历路径上的最小基因均为1，则未出现在该路径上的节点最小路径均为1