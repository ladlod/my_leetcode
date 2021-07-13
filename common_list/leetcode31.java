package common_list;

import java.util.Scanner;

public class leetcode31 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        nextPermutation(nums);
        for (int i = 0; i < n; i++) {
            System.out.print(nums[i] + " ");
        }
        in.close();
    }

    static void nextPermutation(int[] nums) {
        int flag = nums.length - 2;
        while (flag >= 0) {
            if (nums[flag + 1] > nums[flag]) {
                break;
            }
            flag--;
        }
        if (flag >= 0) {
            int tmp = nums[flag];
            int tmp2 = flag + 1;
            for (int i = flag + 1; i < nums.length; i++) {
                if (nums[i] > tmp) {
                    nums[flag] = nums[i];
                    tmp2 = i;
                }
            }
            nums[tmp2] = tmp;
        }
        int start = flag + 1;
        int end = nums.length - 1;
        while (start < end) {
            int t = nums[start];
            nums[start] = nums[end];
            nums[end] = t;
            start++;
            end--;
        }
    }
}