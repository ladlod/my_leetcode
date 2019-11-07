import java.util.Scanner;

public class leetcode5 {
    public static void main(String args[]) {
        String str = "";

        Scanner in = new Scanner(System.in);
        str = in.nextLine();

        /*
         * for (int i = 1; i <= 500; i++) { str += 'a'; }
         */
        System.out.println(longestPalindrome(str));
    }

    public static String longestPalindrome(String s) {
        String result = "";
        String tmp = "";
        for (int i = 0; i < s.length(); i++) {
            tmp = tmp + s.toCharArray()[i];
            int j = 1;
            for (; i - j >= 0 && i + j < s.length(); j++) {
                if (s.toCharArray()[i - j] != s.toCharArray()[i + j])
                    break;
                tmp = s.toCharArray()[i - j] + tmp + s.toCharArray()[i + j];
            }
            if (tmp.length() >= result.length()) {
                result = tmp;
            }
            if (i + j == s.length() - 1) {
                tmp = "";
                break;
            }
            tmp = "";
        }
        for (int i = 0; i < s.length() - 1; i++) {
            int j = 1;
            if (s.toCharArray()[i] == s.toCharArray()[i + 1]) {
                tmp = tmp + s.toCharArray()[i] + s.toCharArray()[i + 1];
                for (; i - j >= 0 && i + j + 1 < s.length(); j++) {
                    if (s.toCharArray()[i - j] != s.toCharArray()[i + 1 + j])
                        break;
                    tmp = s.toCharArray()[i - j] + tmp + s.toCharArray()[i + 1 + j];
                }
                if (tmp.length() >= result.length()) {
                    result = tmp;
                }
                if (i + j + 1 == s.length() - 1) {
                    tmp = "";
                    break;
                }
                tmp = "";
            }
        }
        return result;
    }
}