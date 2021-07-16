# !/usr/bin/python
# -*-coding: utf-8 -*-
def maxProfit(values):
    if len(values) < 2: return 0
    res = 0
    for i in range(1, len(values)):
        if values[i] > values[i-1]: res += values[i] - values[i-1]

    return res

if __name__ == '__main__':
    print(maxProfit([7,6,4,3,1]))