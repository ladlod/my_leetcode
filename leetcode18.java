import java.util.Scanner;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class leetcode18 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n, target;
        target = in.nextInt();
        in.nextLine();
        n = in.nextInt();
        in.nextLine();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        System.out.println(fourSum(nums, target));
        in.close();
    }

    static List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> foursums = new ArrayList<>();
        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1])
                continue;
            for (int j = i + 1; j < nums.length - 2; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1])
                    continue;
                int start = j + 1;
                int end = nums.length - 1;
                while (start < end) {
                    if (nums[i] + nums[j] + nums[start] + nums[end] > target) {
                        end--;
                    } else if (nums[i] + nums[j] + nums[start] + nums[end] < target) {
                        start++;
                    } else {
                        List<Integer> sum = new ArrayList<>();
                        sum.add(nums[i]);
                        sum.add(nums[j]);
                        sum.add(nums[start]);
                        sum.add(nums[end]);
                        foursums.add(sum);
                        while (start < end && nums[start] == nums[start + 1])
                            start++;
                        while (start < end && nums[end] == nums[end - 1])
                            end--;
                        start++;
                        end--;
                    }
                }
            }
        }
        return foursums;
    }
}