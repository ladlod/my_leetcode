package week_contest

// 用一个下标从 0 开始的二维整数数组 rectangles 来表示 n 个矩形，其中 rectangles[i] = [widthi, heighti] 表示第 i 个矩形的宽度和高度。

// 如果两个矩形 i 和 j（i < j）的宽高比相同，则认为这两个矩形 可互换 。更规范的说法是，两个矩形满足 widthi/heighti == widthj/heightj（使用实数除法而非整数除法），则认为这两个矩形 可互换 。

// 计算并返回 rectangles 中有多少对 可互换 矩形。

func interchangeableRectangles(rectangles [][]int) int64 {
    mp := make(map[float64]int)
    for _, rect := range rectangles {
        mp[float64(rect[1])/float64(rect[0])]++
    }
    res := int64(0)
    for _, v := range mp {
        if v >= 2 {
            res += int64(v * (v-1) / 2)
        } 
    }
    return res
}