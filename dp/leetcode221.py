# !/usr/bin/python
# -*-coding: utf-8 -*-

# 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。


def maximalSquare(matrix):
    m = len(matrix)
    if m <= 0: return 0
    n = len(matrix[0])
    square = [[0 for i in range(0, n)] for i in range(0, m)]

    for i in range(0, m):
        for j in range(0, n):
            if i == 0 or j == 0: square[i][j] = 1 if matrix[i][j] == "1" else 0
            else : square[i][j] = min(square[i-1][j-1], square[i][j-1], square[i-1][j]) + 1 if matrix[i][j] == "1" else 0
    
    return max([i for s in square for i in s])**2

if __name__ == '__main__':
    print(maximalSquare([[1,1,1,1,1,1],[1,1,1,1,1,1],[1,1,1,1,1,1]]))