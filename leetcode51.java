import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;

public class leetcode51 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        System.out.println(solveNQueens(n).size());
        System.out.println(solveNQueens(n));
        in.close();
    }

    static List<List<String>> solveNQueens(int n) {
        List<List<String>> res = new ArrayList<>();
        char[][] queen = new char[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                queen[i][j] = '.';
            }
        }
        boolean[] colFlag = new boolean[n];
        boolean[] crossFlag1 = new boolean[2 * n - 1];
        boolean[] crossFlag2 = new boolean[2 * n - 1];
        putQueens(res, queen, 0, colFlag, crossFlag1, crossFlag2);
        return res;
    }

    static void putQueens(List<List<String>> res, char queen[][], int row, boolean[] colFlag, boolean[] crossFlag1,
            boolean[] crossFlag2) {
        int n = queen.length;
        for (int i = 0; i < n; i++) {
            if (colFlag[i] == false && crossFlag1[row - i + n - 1] == false && crossFlag2[row + i] == false) {
                colFlag[i] = true;
                crossFlag1[row - i + n - 1] = true;
                crossFlag2[row + i] = true;
                queen[row][i] = 'Q';
                if (row + 1 == n) {
                    List<String> tmp = new ArrayList<>();
                    for (int j = 0; j < n; j++)
                        tmp.add(String.valueOf(queen[j]));
                    res.add(tmp);
                    colFlag[i] = false;
                    crossFlag1[row - i + n - 1] = false;
                    crossFlag2[row + i] = false;
                    queen[row][i] = '.';
                } else {
                    putQueens(res, queen, row + 1, colFlag, crossFlag1, crossFlag2);
                    colFlag[i] = false;
                    crossFlag1[row - i + n - 1] = false;
                    crossFlag2[row + i] = false;
                    queen[row][i] = '.';
                }
            }
        }
    }
}