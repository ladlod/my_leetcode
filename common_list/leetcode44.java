package common_list;

import java.util.Scanner;

public class leetcode44 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        String p = in.nextLine();
        System.out.println(isMatch(s, p));
        in.close();
    }

    static boolean isMatch(String s, String p) {
        if (s.length() != 0 && p.length() == 0)
            return false;
        else if (s.length() == 0 && p.length() == 0)
            return true;
        boolean[][] res = new boolean[s.length() + 1][p.length() + 1];
        res[0][0] = true;
        for (int i = 0; i < p.length() && p.charAt(i) == '*'; i++) {
            res[0][i + 1] = true;
        }
        for (int i = 0; i < s.length(); i++) {
            for (int j = 0; j < p.length(); j++) {
                if (s.charAt(i) == p.charAt(j) || p.charAt(j) == '?')
                    res[i + 1][j + 1] = res[i][j];
                else if (p.charAt(j) == '*') {
                    if (j == 0)
                        res[i + 1][j + 1] = true;
                    else
                        res[i + 1][j + 1] = res[i + 1][j] || res[i][j + 1];
                }
            }
        }

        for (int i = 0; i <= s.length(); i++) {
            for (int j = 0; j <= p.length(); j++) {
                System.out.print(res[i][j] + " ");
            }
            System.out.println();
        }

        return res[s.length()][p.length()];
    }
}