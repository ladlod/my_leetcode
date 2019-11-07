import java.util.Scanner;

public class leetcode65 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        System.out.println(isNumber(s));
        in.close();
    }

    static boolean isNumber(String s) {
        if (s.length() == 0)
            return false;
        int j = s.length();
        for (; j > 0; j--) {
            if (s.charAt(j - 1) != ' ')
                break;
        }
        s = s.substring(0, j);
        boolean res = false;
        boolean space = true;
        boolean flag = true;
        boolean e = false;
        int e_num = 0;
        int doc_num = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == ' ' && space == true)
                continue;
            else if (s.charAt(i) == ' ' && space == false)
                return false;
            else if (s.charAt(i) == '+' || s.charAt(i) == '-') {
                if (flag == true)
                    flag = false;
                else
                    return false;
            } else if (s.charAt(i) == 'e') {
                if (e == true && e_num == 0) {
                    e_num++;
                    e = false;
                    flag = true;
                    doc_num = 0;
                } else
                    return false;
            } else if (s.charAt(i) == '.') {
                if (doc_num == 0) {
                    doc_num++;
                    space = false;
                } else
                    return false;
            } else if (s.charAt(i) >= '0' && s.charAt(i) <= '9') {
                res = true;
                space = false;
                flag = false;
                e = true;
            } else
                return false;
        }
        if (e == false && e_num == 1)
            return false;

        return res;
    }
}