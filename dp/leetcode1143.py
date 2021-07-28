# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

# 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

# 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
# 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/longest-common-subsequence
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def longestCommonSubsequence(text1, text2):
    l1 = len(text1)
    l2 = len(text2)
    if l1 <= 0 or l2 <= 0: return 0

    res = [[0 for i in range(l2)] for i in range(l1)]
    for i in range(l1):
        for j in range(l2):
            if text1[i] == text2[j] : res[i][j] = (0 if i == 0 or j == 0 else res[i-1][j-1]) + 1
            else : res[i][j] = max(0 if i == 0 else res[i-1][j],0 if j == 0 else res[i][j-1])

    return res[l1-1][l2-1]

if __name__ == '__main__':
    print(longestCommonSubsequence('abcde', 'abce'))