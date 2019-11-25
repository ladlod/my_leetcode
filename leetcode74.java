import java.util.Scanner;

public class leetcode74 {
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
        int target = in.nextInt();
        System.out.println(searchMatrix(matrix, target));
        in.close();
    }

    static boolean searchMatrix(int[][] matrix, int target) {
        for (int i = 0; i < matrix.length; i++) {
            int n = matrix[i].length;
            if (n != 0 && target <= matrix[i][n - 1]) {
                for (int j = 0; j < n; j++) {
                    if (matrix[i][j] == target)
                        return true;
                }
            }
        }
        return false;
    }
}