import java.util.Scanner;

public class leetcode26 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        int len = removeDuplicates(nums);
        System.out.println("len=" + len);
        for (int i = 0; i < len; i++) {
            System.out.print(nums[i] + " ");
        }
        in.close();
    }

    static int removeDuplicates(int[] nums) {
        int len = nums.length;
        int i = 0;
        for (int j = 1; j < len; j++) {
            if (nums[i] != nums[j]) {
                i++;
                nums[i] = nums[j];
            }
        }
        return i + 1;
    }
}