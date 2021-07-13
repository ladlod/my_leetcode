package common_list;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class leetcode38 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        System.out.println(countAndSay(n));
        in.close();
    }

    static String countAndSay(int n) {
        List<String> res = new ArrayList<>();
        res.add("1");
        for (int i = 1; i < n; i++) {
            String pre = res.get(i - 1);
            String s = "";
            int num = 1;
            String tmp = pre.substring(0, 1);
            for (int j = 1; j < pre.length(); j++) {
                if (pre.substring(j, j + 1).equals(tmp))
                    num++;
                else {
                    s = s + (char) (num + '0') + tmp;
                    tmp = pre.substring(j, j + 1);
                    num = 1;
                }
            }
            s = s + (char) (num + '0') + tmp;
            res.add(s);
        }
        return res.get(n - 1);
    }
}