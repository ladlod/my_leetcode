# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和： 

# i - k <= r <= i + k,
# j - k <= c <= j + k 且
# (r, c) 在矩阵内。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/matrix-block-sum
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def matrixBlockSum(mat, k):
    m = len(mat)
    if m <= 0 : return []
    n = len(mat[0])
    square = [[0 for i in range(0, n)] for i in range(0, m)]
    for i in range(0, m):
        for j in range(0, n):
            x = square[i-1][j] if i > 0 else 0
            y = square[i][j-1] if j > 0 else 0
            z = square[i-1][j-1] if i > 0 and j > 0 else 0
            square[i][j] = x + y - z + mat[i][j]
    
    res = [[0 for i in range(0, n)] for i in range(0, m)]
    for i in range(0, m):
        for j in range(0, n):
            l2 = i+k if i+k<m else m-1
            h2 = j+k if j+k<n else n-1
            x = square[i-k-1][h2] if i-k-1 >= 0 else 0
            y = square[l2][j-k-1] if j-k-1 >= 0 else 0
            z = square[i-k-1][j-k-1] if i-k-1 >= 0 and j-k-1 >= 0 else 0

            res[i][j] = square[l2][h2] - x - y + z 

    return res

if __name__ == '__main__':
    print(matrixBlockSum([[1,2,3],[4,5,6],[7,8,9]],1))