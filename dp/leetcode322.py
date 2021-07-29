# !/usr/bin/python
# -*-coding: utf-8 -*-

# 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

# 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

# 你可以认为每种硬币的数量是无限的

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/coin-change
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

def coinChange(coins, amount):
    res = [-1 for i in range(amount+1)]
    res[0] = 0
    for i in range(amount+1):
        for coin in coins:
            if i - coin >= 0:
                if res[i] != -1:
                    if res[i-coin] != -1: res[i] = min(res[i], res[i-coin] + 1)
                    else: res[i] = res[i]
                elif res[i-coin] != -1: res[i] = res[i-coin] + 1

    return res

if __name__ == "__main__":
    print(coinChange([2,3], 10))