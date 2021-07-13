package common_list;

import java.util.Scanner;

public class leetcode34 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        int target = in.nextInt();
        int[] p = searchRange(nums, target);
        System.out.println(p[0] + " " + p[1]);
        in.close();
    }

    static int[] searchRange(int[] nums, int target) {
        int flag = search(nums, 0, nums.length - 1, target);
        System.out.println(flag);
        int[] res = { flag, flag };
        if (flag != -1) {
            while (res[0] >= 0 && nums[res[0]] == target)
                res[0]--;
            while (res[1] <= nums.length - 1 && nums[res[1]] == target)
                res[1]++;
        }
        res[0] += 1;
        res[1] -= 1;
        return res;
    }

    static int search(int[] nums, int begin, int end, int target) {
        if (begin > end)
            return -1;
        else if (begin == end && nums[begin] != target)
            return -1;
        int middle = (begin + end) / 2;
        if (nums[middle] == target)
            return middle;
        else if (target < nums[middle])
            return search(nums, begin, middle - 1, target);
        else
            return search(nums, middle + 1, end, target);
    }
}