package common_list;

import java.util.Scanner;

public class leetcode27 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        int val = in.nextInt();
        int len = removeElement(nums, val);
        System.out.println("len=" + len);
        for (int i = 0; i < len; i++) {
            System.out.print(nums[i] + " ");
        }
        in.close();
    }

    static int removeElement(int[] nums, int val) {
        int len = nums.length;
        int i = 0;
        for (int j = 0; j < len; j++) {
            if (nums[j] != val) {
                nums[i] = nums[j];
                i++;
            }
        }
        return i;
    }
}