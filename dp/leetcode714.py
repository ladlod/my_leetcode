# !/usr/bin/python
# -*-coding: utf-8 -*-

def maxProfit(prices, fee):
    if len(prices) < 2: return 0
    holding = -prices[0]
    sale = 0
    for i in range(1, len(prices)):
        holding, sale = max(holding, sale - prices[i]), max(holding + prices[i] - fee, sale)

    return sale

if __name__ == '__main__':
    print(maxProfit([1,3,7,5,10,3], 3))