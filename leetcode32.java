import java.util.Scanner;

public class leetcode32 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        System.out.println(longestValidParentheses(s));
        in.close();
    }

    static int longestValidParentheses(String s) {
        int res = 0;
        for (int i = 0; i < s.length(); i++) {
            int lmax = 0;
            int len = 0;
            int tag = 0;
            for (int j = i; j < s.length(); j++) {
                len++;
                if (s.charAt(j) == '(')
                    tag++;
                else if (s.charAt(j) == ')') {
                    tag--;
                }
                if (tag < 0)
                    break;
                else if (tag == 0)
                    lmax = len;
            }
            // System.out.println(lmax);
            if (lmax > res)
                res = lmax;
            i += lmax;
        }
        return res;
    }
}