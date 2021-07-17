# !/usr/bin/python
# -*-coding: utf-8 -*-

def maxProfit(prices):
    res1 = -prices[0]
    res2 = 0
    res3 = 0

    for i in range(1, len(prices)):
        res1 = max(res1, res2-prices[i])
        res2 = max(res2, res3)
        res3 = res1 + prices[i]

        return max(res1, res2, res3)

if __name__ == '__main__':
    print(maxProfit([1,2,3,0,2]))