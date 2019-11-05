import java.util.Scanner;
import java.util.Stack;
import java.util.List;
import java.util.ArrayList;

public class leetcode37 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        char[][] board = new char[9][9];
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                board[i][j] = in.next().charAt(0);
            }
        }
        solveSudoku(board);
        System.out.println();
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                System.out.print(board[i][j] + " ");
            }
            System.out.println();
        }
        in.close();
    }

    static void solveSudoku(char[][] board) {
        boolean[][] rowFlag = new boolean[9][9];
        boolean[][] colFlag = new boolean[9][9];
        boolean[][] boxFlag = new boolean[9][9];
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                rowFlag[i][j] = true;
                colFlag[i][j] = true;
                boxFlag[i][j] = true;
            }
        }
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                int x = j / 3 + i / 3 * 3;
                int y = i % 3 * 3 + j % 3;
                if (board[i][j] > '0' && board[i][j] <= '9') {
                    int row = board[i][j] - '0';
                    rowFlag[i][row - 1] = false;
                }
                if (board[i][j] > '0' && board[i][j] <= '9') {
                    int col = board[i][j] - '0';
                    colFlag[j][col - 1] = false;
                }
                if (board[x][y] > '0' && board[x][y] <= '9') {
                    int box = board[x][y] - '0';
                    boxFlag[i][box - 1] = false;
                }
            }
        }
        System.out.println(solveSudo(board, rowFlag, colFlag, boxFlag, 0, 0));
    }

    static boolean solveSudo(char[][] board, boolean[][] rowFlag, boolean[][] colFlag, boolean[][] boxFlag, int row,
            int col) {
        int x = row, y = col;
        b: for (; x < 9; x++) {
            for (; y < 9; y++) {
                if (board[x][y] == ',') {
                    break b;
                }
            }
            y = 0;
        }
        if (x == 9 || y == 9)
            return true;
        Stack<Integer> flag = new Stack<>();
        int boxnum = x / 3 * 3 + y / 3;
        for (int i = 0; i < 9; i++) {
            if (rowFlag[x][i] == true && colFlag[y][i] == true && boxFlag[boxnum][i] == true) {
                flag.add(i + 1);
            }
        }
        if (flag.empty())
            return false;
        while (!flag.empty()) {
            int num = flag.pop();
            board[x][y] = (char) (num + '0');
            rowFlag[x][num - 1] = false;
            colFlag[y][num - 1] = false;
            boxFlag[boxnum][num - 1] = false;
            if (solveSudo(board, rowFlag, colFlag, boxFlag, x, y)) {
                return true;
            } else {
                board[x][y] = ',';
                rowFlag[x][num - 1] = true;
                colFlag[y][num - 1] = true;
                boxFlag[boxnum][num - 1] = true;
            }
        }

        return false;
    }
}