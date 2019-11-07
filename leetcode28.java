import java.util.Scanner;

public class leetcode28 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String haystack = in.nextLine();
        String needle = in.nextLine();
        System.out.println(strStr(haystack, needle));
        in.close();
    }

    static int strStr(String haystack, String needle) {
        int[] next = nextNums(needle);
        int i = 0, j = 0;
        while (i < haystack.length() && j < needle.length()) {
            if (j == -1 || haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
            } else {
                j = next[j];
            }
        }
        if (j == needle.length())
            return i - j;
        else
            return -1;
    }

    static int[] nextNums(String needle) {
        int[] next = new int[needle.length()];
        next[0] = -1;
        int k = -1;
        for (int i = 0; i < needle.length() - 1;) {
            if (k == -1 || needle.charAt(k) == needle.charAt(i)) {
                i++;
                k++;
                next[i] = k;
            } else
                k = next[k];
        }
        /*
         * for (int i = 0; i < needle.length(); i++) { System.out.print(next[i] + " ");
         * } System.out.println();
         */
        return next;
    }
}