package common_list;

import java.util.Scanner;

public class leetcode36 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        char[][] board = new char[9][9];
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                board[i][j] = in.next().charAt(0);
            }
        }
        System.out.println(isValidSudoku(board));
        in.close();
    }

    static boolean isValidSudoku(char[][] board) {
        for (int i = 0; i < 9; i++) {
            int[] flag0 = { 0, 0, 0, 0, 0, 0, 0, 0, 0 };
            int[] flag1 = { 0, 0, 0, 0, 0, 0, 0, 0, 0 };
            int[] flag2 = { 0, 0, 0, 0, 0, 0, 0, 0, 0 };
            for (int j = 0; j < 9; j++) {
                if (board[i][j] >= '0' && board[i][j] <= '9') {
                    int p = board[i][j] - '0';
                    flag0[p - 1]++;
                }
                if (board[j][i] >= '0' && board[j][i] <= '9') {
                    int p = board[j][i] - '0';
                    flag1[p - 1]++;
                }
                int x = j / 3 + i / 3 * 3;
                int y = i % 3 * 3 + j % 3;
                if (board[x][y] >= '0' && board[x][y] <= '9') {
                    int p = board[x][y] - '0';
                    flag2[p - 1]++;
                }
            }
            for (int j = 0; j < 9; j++) {
                if (flag0[j] > 1)
                    return false;
            }
            for (int j = 0; j < 9; j++) {
                if (flag1[j] > 1)
                    return false;
            }
            for (int j = 0; j < 9; j++) {
                if (flag2[j] > 1)
                    return false;
            }
        }
        return true;
    }
}