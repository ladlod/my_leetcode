package dp;

import java.util.Scanner;

public class leetcode42 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] height = new int[n];
        for (int i = 0; i < n; i++) {
            height[i] = in.nextInt();
        }
        System.out.println(trap(height));
        in.close();
    }

    static int trap(int[] height) {
        if (height.length == 0)
            return 0;
        int res = 0;
        int left = 0;
        int right = height.length - 1;
        int lmax = height[left];
        int rmax = height[right];
        while (left < right) {
            if (height[left] < height[right]) {
                if (height[left] <= lmax)
                    res = res + lmax - height[left];
                else
                    lmax = height[left];
                left++;
            } else {
                if (height[right] <= rmax)
                    res = res + rmax - height[right];
                else
                    rmax = height[right];
                right--;
            }
        }
        return res;
    }
}