package week_contest

// 有一棵根节点为 0 的 家族树 ，总共包含 n 个节点，节点编号为 0 到 n - 1 。给你一个下标从 0 开始的整数数组 parents ，其中 parents[i] 是节点 i 的父节点。由于节点 0 是 根 ，所以 parents[0] == -1 。

// 总共有 105 个基因值，每个基因值都用 闭区间 [1, 105] 中的一个整数表示。给你一个下标从 0 开始的整数数组 nums ，其中 nums[i] 是节点 i 的基因值，且基因值 互不相同 。

// 请你返回一个数组 ans ，长度为 n ，其中 ans[i] 是以节点 i 为根的子树内 缺失 的 最小 基因值。

// 节点 x 为根的 子树 包含节点 x 和它所有的 后代 节点。


// 思路
// 暴力解法，遍历所有节点，记录所有节点及其子树出现过的基因，对每个节点的序列排序并求最小未出现的基因，时间复杂路O(n^2*logn)，超时
// 出现在一个节点的值，仅会出现在它及它到根节点的路径上，记录每条路径上出现过的最小数字即可
func smallestMissingValueSubtree(parents []int, nums []int) []int {
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
