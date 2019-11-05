import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.List;

public class leetcode39 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] candidates = new int[n];
        for (int i = 0; i < n; i++) {
            candidates[i] = in.nextInt();
        }
        int target = in.nextInt();
        System.out.println(combinationSum(candidates, target));
        in.close();
    }

    static List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        if (candidates.length == 0)
            return res;
        Arrays.sort(candidates);
        List<Integer> l = new ArrayList<>();
        findCombinationSum(candidates, 0, res, l, target);
        return res;
    }

    static void findCombinationSum(int[] candidates, int start, List<List<Integer>> res, List<Integer> l, int target) {
        int i = start;
        while (i < candidates.length && candidates[i] <= target) {
            int t = target;
            List<Integer> tmp = new ArrayList<>();
            tmp.addAll(l);
            tmp.add(candidates[i]);
            t -= candidates[i];
            if (t == 0)
                res.add(tmp);
            else if (t > 0)
                findCombinationSum(candidates, i, res, tmp, t);
            i++;
        }
    }
}