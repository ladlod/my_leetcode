package common_list;

import java.util.Scanner;

public class leetcode80{
    public static void main(String[] args){ 
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for(int i = 0; i < n; i++){
            nums[i] = in.nextInt();
        }
        System.out.println(removeDuplicates(nums));
        for(int i = 0; i < n; i++){
            System.out.print(nums[i] + " ");
        }
        System.out.println();
        in.close();
    }

    static int removeDuplicates(int[] nums) {
        int x = 2;
        for(int y = 2; y < nums.length; y++){
            nums[x] = nums[y];
            if(!(nums[x] == nums[x - 1] && nums[x] == nums[x - 2])){
                x++;
            }
        }

        return x;
    }
}