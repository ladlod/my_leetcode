import java.util.Scanner;

public class leetcode43 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String num1 = in.nextLine();
        String num2 = in.nextLine();
        System.out.println(multiply(num1, num2));
        in.close();
    }

    static String multiply(String num1, String num2) {
        char[] mul = new char[num1.length() + num2.length()];
        for (int i = 0; i < mul.length; i++) {
            mul[i] = '0';
        }
        for (int i = num1.length() - 1; i >= 0; i--) {
            for (int j = num2.length() - 1; j >= 0; j--) {
                int add = mul[i + j + 1] - '0' + (num1.charAt(i) - '0') * (num2.charAt(j) - '0');
                mul[i + j + 1] = (char) (add % 10 + '0');
                mul[i + j] += (char) (add / 10);
            }
        }
        String res = "";
        int start = 0;
        for (; start < mul.length; start++) {
            if (mul[start] != '0')
                break;
        }
        if (start == mul.length)
            return "0";
        for (; start < mul.length; start++) {
            res += mul[start];
        }
        return res;
    }
}