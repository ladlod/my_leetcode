# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

# 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

# 假设每一种面额的硬币有无限个。 

# 题目数据保证结果符合 32 位带符号整数。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/coin-change-2
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def change(amount, coins):
    res = [0 for i in range(amount+1)]
    res[0] = 1
    for coin in coins:
        for i in range(coin, amount + 1):
            res[i] += res[i-coin]

    return res

if __name__ == "__main__":
    print(change(10, [2,3]))