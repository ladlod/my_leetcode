# !/usr/bin/python
# -*-coding: utf-8 -*-
def maxProfit1(values):
    if len(values) < 2: return 0
    income = values[1] - values[0]
    res = income
    for i in range(2, len(values)):
        income = values[i] - values[i-1] + (income if income > 0 else 0)
        res = max(income, res)
    return max(res, 0)

if __name__ == '__main__':
    print(maxProfit1([7,1,5,3,6,4]))