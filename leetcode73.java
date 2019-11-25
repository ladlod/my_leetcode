import java.util.Scanner;

public class leetcode73 {
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
        setZeroes(matrix);
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                System.out.print(matrix[i][j] + " ");
            }
            System.out.println();
        }
        in.close();
    }

    static void setZeroes(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        boolean row0 = true;
        boolean col0 = true;
        for (int i = 0; i < m; i++) {
            if (matrix[i][0] == 0) {
                row0 = false;
                break;
            }
        }
        for (int j = 0; j < n; j++) {
            if (matrix[0][j] == 0) {
                col0 = false;
                break;
            }
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (matrix[i][j] == 0) {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }
        for (int i = 1; i < m; i++) {
            if (matrix[i][0] == 0) {
                for (int j = 1; j < n; j++) {
                    matrix[i][j] = 0;
                }
            }
        }
        for (int j = 1; j < n; j++) {
            if (matrix[0][j] == 0) {
                for (int i = 1; i < m; i++) {
                    matrix[i][j] = 0;
                }
            }
        }
        if (!row0) {
            for (int i = 0; i < m; i++)
                matrix[i][0] = 0;
        }
        if (!col0) {
            for (int j = 0; j < n; j++)
                matrix[0][j] = 0;
        }
    }
}