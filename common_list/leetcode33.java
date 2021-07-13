package common_list;

import java.util.Scanner;

public class leetcode33 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        int target = in.nextInt();
        System.out.println(search(nums, target));
        in.close();
    }

    static int search(int[] nums, int target) {
        if (nums.length == 0)
            return -1;
        if (nums.length == 1)
            return nums[0] == target ? 0 : -1;
        int flag = searchC(nums, 0, nums.length - 1);
        if (target == nums[0])
            return 0;
        else if (target > nums[0])
            return searchW(nums, 1, flag, target);
        else
            return searchW(nums, flag + 1, nums.length - 1, target);
    }

    static int searchC(int[] nums, int begin, int end) {
        if (begin == end)
            return nums.length - 1;
        int middle = (begin + end) / 2;
        if (nums[middle] > nums[middle + 1])
            return middle;
        if (nums[middle] < nums[begin])
            return searchC(nums, begin, middle);
        else if (nums[middle] > nums[end])
            return searchC(nums, middle + 1, end);
        else
            return nums.length - 1;
    }

    static int searchW(int[] nums, int begin, int end, int target) {
        if (begin > end)
            return -1;
        else if (begin == end && nums[begin] != target)
            return -1;
        int middle = (begin + end) / 2;
        if (nums[middle] == target)
            return middle + 1;
        else if (target < nums[middle])
            return searchW(nums, begin, middle - 1, target);
        else
            return searchW(nums, middle + 1, end, target);
    }
}