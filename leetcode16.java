import java.util.Scanner;
import java.util.Arrays;

public class leetcode16 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n, target;
        target = in.nextInt();
        n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        System.out.println(threeSum(nums, target));
        in.close();
    }

    public static int threeSum(int[] nums, int target) {
        int ret = nums[0] + nums[1] + nums[2];
        Arrays.sort(nums);
        for (int i = 0; i < nums.length; i++) {
            if (i > 0 && nums[i] == nums[i - 1])
                continue;
            int start = i + 1;
            int end = nums.length - 1;
            while (start < end) {
                if (nums[i] + nums[start] + nums[end] > target) {
                    int _ret = ret < target ? target - ret : ret - target;
                    int _sum = nums[i] + nums[start] + nums[end] - target;
                    if (_sum < _ret)
                        ret = nums[i] + nums[start] + nums[end];
                    end--;
                } else if (nums[i] + nums[start] + nums[end] < target) {
                    int _ret = ret < target ? target - ret : ret - target;
                    int _sum = target - nums[i] - nums[start] - nums[end];
                    if (_sum < _ret)
                        ret = nums[i] + nums[start] + nums[end];
                    start++;
                } else {
                    return target;
                }
            }
        }
        return ret;
    }
}