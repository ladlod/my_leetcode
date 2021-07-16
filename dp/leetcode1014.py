# !/usr/bin/python
# -*-coding: utf-8 -*-
def maxScoreSightseeingPair(values):
    if len(values) < 2: return 0
    score = values[0] + values[1] - 1
    res = score
    for i in range(2, len(values)):
        score = max(values[i] + score - values[i-1] - 1, values[i] + values[i-1] -1)
        res = max(score, res)

    return res

if __name__ == '__main__':
    print(maxScoreSightseeingPair([8,1,5,2,6]))