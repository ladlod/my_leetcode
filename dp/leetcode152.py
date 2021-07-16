# !/usr/bin/python
# -*-coding: utf-8 -*-

def maxProduct(nums):
        if len(nums) <= 0:
                return 0
        max = nums[0]
        min = nums[0]
        res = nums[0]
        for i in range(1, len(nums)):
                if nums[i] > 0 :
                        max, min = (max*nums[i] if max > 0 else nums[i]), (min*nums[i] if min < 0 else nums[i])
                else:
                        max, min = (min*nums[i] if min < 0 else nums[i]), (max*nums[i] if max > 0 else nums[i])
                if max > res:
                        res = max
                
        return res

if __name__ == '__main__':
        print(maxProduct(nums=[-1,-2,-9,-6]))
