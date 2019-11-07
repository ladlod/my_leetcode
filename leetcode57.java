import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.Stack;

public class leetcode57 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[][] intervals = new int[n][2];
        int[] newInterval = new int[2];
        for (int i = 0; i < n; i++) {
            intervals[i][0] = in.nextInt();
            intervals[i][1] = in.nextInt();
        }
        newInterval[0] = in.nextInt();
        newInterval[1] = in.nextInt();
        int[][] res = insert(intervals, newInterval);
        for (int i = 0; i < res.length; i++) {
            System.out.println(res[i][0] + " " + res[i][1]);
        }
        in.close();
    }

    static int[][] insert(int[][] intervals, int[] newInterval) {
        int i = 0, j = 0;
        List<int[]> l = new ArrayList<>();
        while (j < intervals.length) {
            if (newInterval[0] > intervals[j][1]) {
                i++;
                j++;
            } else if (newInterval[1] < intervals[j][0]) {
                break;
            } else {
                j++;
            }
        }
        System.out.println("i=" + i + " j=" + j);
        for (int x = 0; x < i; x++) {
            l.add(intervals[x]);
        }
        int[] m = new int[2];
        if (i == intervals.length)
            m[0] = newInterval[0];
        else
            m[0] = newInterval[0] < intervals[i][0] ? newInterval[0] : intervals[i][0];
        if (j == 0)
            m[1] = newInterval[1];
        else
            m[1] = newInterval[1] > intervals[j - 1][1] ? newInterval[1] : intervals[j - 1][1];
        l.add(m);
        for (int x = j; x < intervals.length; x++) {
            l.add(intervals[x]);
        }
        int[][] res = new int[l.size()][2];
        for (int x = 0; x < l.size(); x++) {
            res[x] = l.get(x);
        }
        return res;
    }
}