import java.util.Scanner;

public class leetcode29 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int dividend = in.nextInt();
        int divisor = in.nextInt();
        System.out.println(divide(dividend, divisor));
        in.close();
    }

    static int divide(int dividend, int divisor) {
        if (dividend == -2147483648 && divisor == -1)
            return 2147483647;
        int op = 1;
        if (dividend < 0) {
            dividend = -dividend;
            op = -op;
        }
        if (divisor < 0) {
            divisor = -divisor;
            op = -op;
        }
        while (dividend < divisor)
            return 0;
        int ret = 0;
        int tmp = 1;
        int d = divisor;
        while (dividend - d > d) {
            d = d + d;
            tmp = tmp + tmp;
        }
        ret = ret + tmp + divide(dividend - d, divisor);
        if (op < 0)
            ret = -ret;

        return ret;
    }
}