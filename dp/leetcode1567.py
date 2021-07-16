# !/usr/bin/python
# -*-coding: utf-8 -*-

def getMaxLen(nums):
    if len(nums) <= 0: return 0

    pos, neg = (1 if nums[0] > 0 else 0), (1 if nums[0] < 0 else 0)
    res = pos
    for i in range(1, len(nums)):
        if nums[i] == 0: pos = neg = 0
        elif nums[i] > 0:
            pos, neg = (pos + 1 if pos != 0 else 1), (neg + 1 if neg != 0 else 0)
        elif nums[i] < 0:
            pos, neg = (neg + 1 if neg != 0 else 0), (pos + 1 if pos != 0 else 1)
        if pos > res:
            res = pos

    return res

if __name__ == '__main__':
        print(getMaxLen([1]))