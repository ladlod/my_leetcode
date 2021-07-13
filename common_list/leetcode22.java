package common_list;

import java.util.Scanner;
import java.util.ArrayList;
import java.util.List;

public class leetcode22 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        System.out.println(generateParenthesis(n));
        in.close();
    }

    static List<String> generateParenthesis(int n) {
        List<List<String>> strs = new ArrayList<>();
        List<String> str1 = new ArrayList<>();
        List<String> str2 = new ArrayList<>();
        str1.add("");
        str2.add("()");
        strs.add(str1);
        strs.add(str2);
        for (int i = 2; i <= n; i++) {
            List<String> tmp = new ArrayList<>();
            for (int j = 0; j < i; j++) {
                int p = 0;
                while (p < strs.get(j).size()) {
                    int q = 0;
                    while (q < strs.get(i - 1 - j).size()) {
                        String str = new String();
                        if (strs.get(j).size() == 0)
                            str = "()" + strs.get(i - 1 - j).get(q);
                        else
                            str = "(" + strs.get(j).get(p) + ")" + strs.get(i - 1 - j).get(q);
                        q++;
                        tmp.add(str);
                    }
                    p++;
                }
            }
            strs.add(tmp);
            // sSystem.out.println(strs);
        }
        return strs.get(n);
    }
}