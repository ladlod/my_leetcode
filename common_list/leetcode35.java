package common_list;

import java.util.Scanner;

public class leetcode35 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        int target = in.nextInt();
        System.out.println(searchInsert(nums, target));
        in.close();
    }

    static int searchInsert(int[] nums, int target) {
        return search(nums, 0, nums.length - 1, target);
    }

    static int search(int[] nums, int begin, int end, int target) {
        if (end == 0)
            return target <= nums[end] ? 0 : 1;
        else if (begin == nums.length - 1)
            return target > nums[begin] ? nums.length : begin;
        int middle = (begin + end) / 2;
        // System.out.println(middle);
        if (nums[middle] == target)
            return middle;
        else if (target < nums[middle] && (middle == 0 || target > nums[middle - 1]))
            return middle;
        else if (target > nums[middle] && (middle == nums.length - 1 || target < nums[middle + 1]))
            return middle + 1;
        else if (target > nums[middle])
            return search(nums, middle + 1, end, target);
        else
            return search(nums, begin, middle - 1, target);
    }
}