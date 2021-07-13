package common_list;

import java.util.Scanner;

public class leetcode59 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[][] res = generateMatrix(n);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                System.out.print(res[i][j] + " ");
            }
            System.out.println();
        }
        in.close();
    }

    static int[][] generateMatrix(int n) {
        int[][] res = new int[n][n];
        if (n == 0)
            return res;
        int row = 0;
        int value = 1;
        n -= 1;
        while (row <= n) {
            if (row == n) {
                res[n][n] = value;
            } else {
                for (int i = row; i < n; i++) {
                    res[row][i] = value;
                    value++;
                }
                for (int i = row; i < n; i++) {
                    res[i][n] = value;
                    value++;
                }
                for (int i = 0; i < n - row; i++) {
                    res[n][n - i] = value;
                    value++;
                }
                for (int i = 0; i < n - row; i++) {
                    res[n - i][row] = value;
                    value++;
                }
            }

            row++;
            n--;
        }
        return res;
    }
}