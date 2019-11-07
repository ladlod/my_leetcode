import java.util.Arrays;
import java.util.Scanner;
import java.util.Stack;

public class leetcode56 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[][] intervals = new int[n][2];
        for (int i = 0; i < n; i++) {
            intervals[i][0] = in.nextInt();
            intervals[i][1] = in.nextInt();
        }
        int[][] res = merge(intervals);
        for (int i = 0; i < res.length; i++) {
            System.out.println(res[i][0] + " " + res[i][1]);
        }
        in.close();
    }

    static int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> a[0] - b[0]);
        Stack<int[]> s = new Stack<>();
        for (int i = 0; i < intervals.length; i++) {
            if (s.empty()) {
                s.push(intervals[0]);
            } else {
                int[] tmp = s.pop();
                if (intervals[i][0] <= tmp[1] && intervals[i][1] > tmp[1]) {
                    tmp[1] = intervals[i][1];
                    s.push(tmp);
                } else if (intervals[i][0] > tmp[1]) {
                    s.push(tmp);
                    s.push(intervals[i]);
                } else {
                    s.push(tmp);
                }
            }
        }
        int[][] res = new int[s.size()][2];
        int i = s.size() - 1;
        while (!s.empty()) {
            int[] tmp = s.pop();
            res[i] = tmp;
            i--;
        }
        return res;
    }
}