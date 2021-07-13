package dp;
import java.util.Scanner;

public class leetcode45 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        System.out.println(jump(nums));
        in.close();
    }

    static int jump(int[] nums) {
        int step = 0;
        int end = nums.length - 1;
        while (end > 0) {
            for (int i = 0; i <= end; i++) {
                if (i + nums[i] >= end) {
                    end = i;
                    break;
                }
            }
            step++;
            System.out.print(end + " ");
        }
        System.out.println();

        return step;
    }
}