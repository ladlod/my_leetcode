import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;

public class leetcode78{
    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int nums[] = new int[n];
        for(int i = 0; i < n; i++){
            nums[i] = in.nextInt();
        }

        System.out.println(subsets(nums));

        in.close();
    }
    static List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        return my_subsets(nums, res, 0);
    }

    static List<List<Integer>> my_subsets(int[] nums, List<List<Integer>> res, int n){
        if(n == 0){
            res.add(new ArrayList<Integer>());
        }
        else{
            int l = res.size();
            for(int i = 0; i < l; i++){
                List<Integer> tmp = new ArrayList<Integer>(res.get(i));
                tmp.add(nums[n - 1]);
                res.add(tmp);
            }
        }

        System.out.println(res);

        if(n < nums.length)
            return my_subsets(nums, res, n + 1);
        else
            return res;
    }
}