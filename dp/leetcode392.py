# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

# 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

# 进阶：

# 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/is-subsequence
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def isSubsequence(s, t):
    if len(s) == 0 : return True
    elif len(t) == 0 : return False

    res = [[False for i in range(0, len(t))] for i in range(0, len(s))]
    for i in range(0, len(s)):
        for j in range(i, len(t)):
            if s[i] == t[j]: res[i][j] = True if i == 0 else res[i-1][j-1]
            else: res[i][j] = False if j == 0 else res[i][j-1]

    return res

if __name__ == '__main__':
    print(isSubsequence('123','1243'))