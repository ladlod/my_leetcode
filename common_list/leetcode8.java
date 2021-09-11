package common_list;

import java.util.Scanner;

public class leetcode8 {
    public static void main(String atgs[]) {
        Scanner in = new Scanner(System.in);
        String str = in.nextLine();
        System.out.println(myAtoi(str));
        in.close();
    }

    public static int myAtoi(String str) {
        int ret = 0;
        int sign = 1;
        int start = 0;
        int over = 0;
        char[] array = str.toCharArray();
        for (int i = 0; i < str.length(); i++) {
            if ((array[i] >= '0' && array[i] <= '9')) {
                if (ret > Integer.MAX_VALUE / 10 || (ret == Integer.MAX_VALUE / 10 && array[i] >= '8')) {
                    ret = Integer.MAX_VALUE;
                    over = 1;
                    break;
                }
                start = 1;
                ret = ret * 10 + array[i] - '0';
            } else if (array[i] == '-') {

                if (start == 1)
                    break;
                sign = 0;
                start = 1;
            } else if (array[i] == '+') {

                if (start == 1)
                    break;
                sign = 1;
                start = 1;
            } else if (array[i] == ' ') {
                if (start == 1)
                    break;
            } else
                break;

        }
        if (sign == 0) {
            ret = -ret;
            if (over == 1)
                ret -= 1;
        }

        return ret;
    }
}