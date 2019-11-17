import java.util.Scanner;

public class leetcode6 {
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        in.nextLine();
        String s = in.nextLine();
        System.out.println(convert(s, n));
        in.close();
    }

    public static String convert(String s, int numRows) {
        if (numRows == 1 || s.length() <= numRows)
            return s;
        String[] Str = new String[numRows];
        for (int i = 0; i < numRows; i++)
            Str[i] = "";
        int row = 0;
        int j = 0;
        for (int i = 0; i < s.length(); i++) {
            if (i % (2 * numRows - 2) >= numRows)
                row = i % (2 * numRows - 2) - numRows + 1;
            else
                row = 0;
            if (row == 0) {
                Str[j] += s.toCharArray()[i];
                j++;
            } else {
                for (int k = 0; k < numRows; k++) {
                    if (k != numRows - row - 1)
                        Str[k] += " ";
                }
                Str[numRows - row - 1] += s.toCharArray()[i];
            }
            if (j == numRows)
                j = 0;
        }
        String ret = "";
        for (int i = 0; i < numRows; i++) {
            System.out.println(Str[i]);
            ret += Str[i];
        }
        return ret;
    }
}