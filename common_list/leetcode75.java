package common_list;

import java.util.Scanner;

public class leetcode75 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] matrix = new int[n];
        for (int i = 0; i < n; i++) {
            matrix[i] = in.nextInt();
        }
        sortColors(matrix);
        for(int i = 0; i < n; i++)
            System.out.print(matrix[i] + " ");
        in.close();
    }

    static void sortColors(int[] nums) {
        int i = 0, begin = 0, end = nums.length - 1;
        while (i <= end) {
            if(nums[i] == 0){
                int tmp = nums[i];
                nums[i] = nums[begin];
                nums[begin] = tmp;
                begin++;
                i++;
            }
            else if(nums[i] == 2){
                int tmp = nums[i];
                nums[i] = nums[end];
                nums[end] = tmp;
                end--;
            }
            else
                i++;
        }
    }
}