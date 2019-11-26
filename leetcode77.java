import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Scanner;

public class leetcode77 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int k = in.nextInt();
        System.out.println(combine(n, k));
        in.close();
    }

    static List<List<Integer>> combine(int n, int k) {
        Queue<List<Integer>> res = new LinkedList<>();
        res.add(new LinkedList<Integer>());
        int i = 1;
        while(i <= k){
            int l = res.size();
            for(int p = 0; p < l; p++){
                //System.out.println(res);
                List<Integer> tmp = res.remove();
                int j = i;
                if(!tmp.isEmpty())
                    j = tmp.get(tmp.size() - 1) + 1;
                for(; j <= n - k + i; j++){
                    List<Integer> tmpc = new LinkedList<>();
                    tmpc.addAll(tmp);
                    tmpc.add(j);
                    res.add(tmpc);
                }
            }
            i++;
        }
        return new LinkedList<List<Integer>>(res);
    }
}