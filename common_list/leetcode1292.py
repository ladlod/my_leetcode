# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。

# 请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# [1,1,3,2,4,3,2]
# [1,1,3,2,4,3,2]
# [1,1,3,2,4,3,2]

def maxSideLength(mat, threshold):
    m = len(mat)
    if m <= 0 : return 0
    n = len(mat[0])
    square = [[0 for i in range(0, n)] for i in range(0, m)]
    for i in range(0, m):
        for j in range(0, n):
            x = square[i-1][j] if i > 0 else 0
            y = square[i][j-1] if j > 0 else 0
            z = square[i-1][j-1] if i > 0 and j > 0 else 0
            square[i][j] = x + y - z + mat[i][j]

    res = -1
    for l in range(1,min(m,n)):
        for i in range(0,m-l):
            for j in range(0,n-l):
                x = square[i-1][j+l] if i > 0 else 0
                y = square[i+l][j-1] if j > 0 else 0
                z = square[i-1][j-1] if i > 0 and j > 0 else 0

                if square[i+l][j+l] - x - y + z <= threshold:
                    res = max(res, l)
                    continue
    
    return res + 1

if __name__ == '__main__':
    print(maxSideLength([[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]],6))