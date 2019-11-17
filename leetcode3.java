import java.util.Scanner;

public class leetcode3 {
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        String str = in.nextLine();
        int l = lengthOfLongestSubstring(str);
        System.out.println(l);
        in.close();
    }

    static int lengthOfLongestSubstring(String s) {
        String max = "";
        String tmp = "";
        for (int i = 0; i < s.length(); i++) {
            for (int j = 0; j < tmp.length(); j++) {
                if (s.toCharArray()[i] == tmp.toCharArray()[j]) {
                    // System.out.println("sub");
                    if (tmp.length() > max.length())
                        max = tmp;
                    tmp = tmp.substring(j + 1);
                    break;
                }
            }
            tmp = tmp + s.toCharArray()[i];
            // System.out.println(3 / 2);
            if (i == s.length() - 1 && tmp.length() > max.length())
                max = tmp;
        }
        return max.length();
    }
}