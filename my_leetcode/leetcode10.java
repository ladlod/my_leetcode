
/*给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
*/

import java.util.Scanner;

public class leetcode10 {
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        String p = in.nextLine();
        System.out.println(isMatch(s, p));
    }

    public static boolean isMatch(String s, String p) {
        if (s.equals(p))
            return true;
        int slen = s.length();
        int plen = p.length();
        boolean[][] sign = new boolean[slen + 1][plen + 1];
        sign[0][0] = true;

        if (slen == 0) {
            for (int j = 1; j <= plen; j++) {
                if (p.charAt(j - 1) == '*' && j >= 2) {
                    sign[0][j] = sign[0][j - 2];
                }
            }
        }
        for (int i = 1; i <= slen; i++) {
            for (int j = 1; j <= plen; j++) {
                if (s.charAt(i - 1) == p.charAt(j - 1) || p.charAt(j - 1) == '.') {
                    sign[i][j] = sign[i - 1][j - 1];
                } else if (p.charAt(j - 1) == '*') {
                    if (j > 2) {
                        sign[0][j] = sign[0][j - 2];
                        sign[i][j] = sign[i][j - 2]
                                || (sign[i - 1][j] && (s.charAt(i - 1) == p.charAt(j - 2) || p.charAt(j - 2) == '.'));
                    } else {
                        sign[0][j] = true;
                        sign[i][j] = sign[i - 1][j] && (s.charAt(i - 1) == p.charAt(j - 2) || p.charAt(j - 2) == '.');
                    }
                }
            }
        }

        return sign[slen][plen];
    }
}