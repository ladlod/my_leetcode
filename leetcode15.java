import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class leetcode15 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n;
        n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        System.out.println(threeSum(nums));
        in.close();
    }

    public static List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> threesums = new ArrayList<>();
        Arrays.sort(nums);
        for (int i = 0; i < nums.length && nums[i] <= 0; i++) {
            if (i > 0 && nums[i] == nums[i - 1])
                continue;
            int start = i + 1;
            int end = nums.length - 1;
            while (start < end) {
                if (nums[i] + nums[start] + nums[end] > 0)
                    end--;
                else if (nums[i] + nums[start] + nums[end] < 0)
                    start++;
                else {
                    List<Integer> sum = new ArrayList<>();
                    sum.add(nums[i]);
                    sum.add(nums[start]);
                    sum.add(nums[end]);
                    threesums.add(sum);
                    while (start < end && nums[start] == nums[start + 1])
                        start++;
                    while (start < end && nums[end] == nums[end - 1])
                        end--;
                    start++;
                    end--;
                }
            }
        }
        return threesums;
    }
}