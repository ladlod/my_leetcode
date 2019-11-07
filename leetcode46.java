import java.util.Scanner;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Arrays;

public class leetcode46 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = in.nextInt();
        }
        System.out.println(permute(nums));
        in.close();
    }

    static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        if (nums.length == 0)
            return res;
        Arrays.sort(nums);
        Queue<List<Integer>> s = new LinkedList<>();
        List<Integer> l = new LinkedList<>();
        l.add(nums[0]);
        s.add(l);
        for (int i = 1; i < nums.length; i++) {
            int deepth = s.size();
            while (deepth > 0) {
                List<Integer> tmp1 = s.poll();
                for (int j = 0; j <= tmp1.size(); j++) {
                    if (j > 0 && nums[i] == nums[i - 1])
                        continue;
                    List<Integer> tmp2 = new LinkedList<>();
                    tmp2.addAll(tmp1);
                    tmp2.add(j, nums[i]);
                    s.add(tmp2);
                }
                deepth--;
            }
        }
        res.addAll(s);
        return res;
    }
}