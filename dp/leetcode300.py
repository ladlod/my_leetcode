# !/usr/bin/python
# -*-coding: utf-8 -*-
# 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

# 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# 动态规划做法
def lengthOfLISDp(nums):
    l = len(nums)
    res = [1 for i in range(l)]
    for i in range(l):
        for j in range(i):
            if nums[i] > nums[j]: res[i] = max(res[i], res[j] + 1)

    return res

# 贪心做法
def lengthOfLISTx(nums):
    l = len(nums)
    if l <= 0: return 0

    res = []
    res.append(nums[0])
    for num in nums:
        if num > res[len(res)-1]: res.append(num)
        else:
            l = 0
            loc = r = len(res)-1
            while l <= r:
                t = int((l+r)/2)
                if num <= res[t]: 
                    loc = t
                    r = t-1
                elif num > res[t]:
                    l = t+1
            res[loc] = num
        print(res)

    return res
    

if __name__ == "__main__":
    print(lengthOfLISDp([1,3,6,7,9,4,10,5,6]))
    print(lengthOfLISTx([18,55,66,2,3,54]))