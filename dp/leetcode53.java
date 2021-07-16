package dp;

public class leetcode53 {
    public int maxSubArray(int[] nums) {
        if(nums.length == 0)
            return 0;
        int len = nums.length;
        int[] sum = new int[len];
        sum[0] = nums[0];
        for(int i = 1; i < len; i++){
            if(sum[i - 1] >= 0)
                sum[i] = nums[i] + sum[i - 1];
            else
                sum[i] = nums[i];
        }
        int res = sum[0];
        for(int i = 1; i < len; i++){
            if(sum[i] > res)
                res = sum[i];
        }
        return res;
    }
}
