# !/usr/bin/python
# -*-coding: utf-8 -*-

def nthUglyNumber(n):
    res = [1,2,3,4,5]
    p2 = 2
    p3 = 2
    p5 = 4
    while len(res) < n:
        add = min(2*res[p2], 3*res[p3], 5*res[p5])
        if add == 2*res[p2]:
            p2 += 1
        elif add == 3*res[p3]:
            p3+=1
        elif add == 5*res[p5]:
            p5+=1
        if res[len(res)-1] != add:
            res.append(add)

    return res

if __name__ == '__main__':
    print(nthUglyNumber(30))