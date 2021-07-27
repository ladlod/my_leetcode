# !/usr/bin/python
# -*-coding: utf-8 -*-

# 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

# 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。

# 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
# 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。

# 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。



# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/wiggle-subsequence
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def wiggleMaxLength(nums):
    l = len(nums)
    if l <= 0: return 0
    elif l == 1: return 1
    elif l == 2: return 2 if nums[0] != nums[1] else 1
    else:
        sub = []
        for i in range(1,l):
            sub.append(nums[i]-nums[i-1])
        l = len(sub)
        res = [1 if sub[i] != 0 else 0 for i in range(l)]
        for i in range(l):
            for j in range(i):
                if sub[i]*sub[j] < 0: res[i] = max(res[i], res[j] + 1)

        return max(res) + 1

if __name__ == "__main__":
    print(wiggleMaxLength([1, 7, 4, 9, 2, 5]))