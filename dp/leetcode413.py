# !/usr/bin/python
# -*-coding: utf-8 -*-
def numberOfArithmeticSlices(nums):
    if len(nums) < 3:
        return 0

    list = []
    start = 0
    step = nums[1]-nums[0]
    for i in range(2, len(nums)):
        if nums[i] - nums[i-1] != step :
            if i - start > 2:
                list.append(i-start)
            start = i-1
            step = nums[i] - nums[i-1]
    if len(nums) - start > 2:
        list.append(len(nums)-start)

    res = 0
    for v in list:
        res += (1+v-2)*(v-2)/2


    return int(res)

if __name__ == '__main__':
    print(numberOfArithmeticSlices([7, 7, 7, 7]))