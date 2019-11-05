import java.util.Scanner;;

public class leetcode14 {
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int n;
        n = in.nextInt();
        in.nextLine();
        String[] strs = new String[n];
        for (int i = 0; i < n; i++) {
            strs[i] = in.nextLine();
        }
        System.out.println(longestCommonPrefix(strs));
        in.close();
    }

    public static String longestCommonPrefix(String[] strs) {
        if (strs.length == 0)
            return "";
        for (int i = 0; i < strs.length; i++) {
            if (strs[i].length() == 0)
                return "";
        }
        String str = "";
        roll: for (int i = 0;; i++) {
            char w = strs[0].charAt(i);
            for (int j = 1; j < strs.length; j++) {
                if (i == strs[j].length())
                    break roll;
                if (strs[j].charAt(i) != w)
                    break roll;
            }
            str += w;
            if (i == strs[0].length() - 1)
                break;
        }

        return str;
    }
}