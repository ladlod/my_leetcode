import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;

public class leetcode54 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int m = in.nextInt();
        int n = in.nextInt();
        int[][] matrix = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                matrix[i][j] = in.nextInt();
            }
        }
        System.out.println(spiralOrder(matrix));
        in.close();
    }

    static List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        int m = matrix.length - 1;
        if (m == -1)
            return res;
        int n = matrix[0].length - 1;
        int row = 0;
        int col = 0;
        while (row <= m && col <= n) {
            if (row == m && col == n) {
                res.add(matrix[m][n]);
            } else if (row == m) {
                for (int i = col; i <= n; i++) {
                    res.add(matrix[row][i]);
                }
            } else if (col == n) {
                for (int i = row; i <= m; i++) {
                    res.add(matrix[i][n]);
                }
            } else {
                for (int i = col; i < n; i++) {
                    res.add(matrix[row][i]);
                }
                for (int i = row; i < m; i++) {
                    res.add(matrix[i][n]);
                }
                for (int i = 0; i < n - col; i++) {
                    res.add(matrix[m][n - i]);
                }
                for (int i = 0; i < m - row; i++) {
                    res.add(matrix[m - i][col]);
                }
            }

            row++;
            col++;
            m--;
            n--;
        }
        return res;
    }
}