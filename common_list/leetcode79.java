package common_list;

import java.util.Scanner;

public class leetcode79{
    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        char[][] words = {{'A', 'B', 'C', 'E'},
                          {'S', 'F', 'C', 'S'},
                          {'A', 'D', 'E', 'E'}};
        String word = in.nextLine();

        System.out.println(exist(words, word));

        in.close();
    }

    static boolean exist(char[][] board, String word) {
        if(board.length == 0 && board[0].length == 0)
            return false;
        int lenx = board.length, leny = board[0].length;
        boolean[][] judge = new boolean[lenx][leny];
        for(int i = 0; i < lenx; i++){
            for(int j = 0; j < leny; j++){
                judge[i][j] = true;
            }
        }
        for(int x = 0; x < lenx; x++){
            for(int y = 0; y < leny; y++){
                if(board[x][y] == word.charAt(0)){
                    judge[x][y] = false;
                    if(my_exist(board, word, x, y, 0, judge))
                        return true;
                    judge[x][y] = true;
                }
            }
        }
        return false;
    }

    static boolean my_exist(char[][] board, String word, int x, int y, int n, boolean[][] judge){
        if(n == word.length() - 1){
            return true;
        }
        if(x + 1 < board.length && judge[x + 1][y] && board[x + 1][y] == word.charAt(n + 1)){
            judge[x + 1][y] = false;
            if(my_exist(board, word, x + 1, y, n + 1, judge))
                return true;
            judge[x + 1][y] = true;
        }

        if(y + 1 < board[0].length && judge[x][y + 1] && board[x][y + 1] == word.charAt(n + 1)){
            judge[x][y + 1] = false;
            if(my_exist(board, word, x, y + 1, n + 1, judge))
                return true;
            judge[x][y + 1] = true;
        }

        if(x > 0 && judge[x - 1][y] && board[x - 1][y] == word.charAt(n + 1)){
            judge[x - 1][y] = false;
            if(my_exist(board, word, x - 1, y, n + 1, judge))
                return true;
            judge[x - 1][y] = true;
        }

        if(y > 0 && judge[x][y - 1] && board[x][y - 1] == word.charAt(n + 1)){
            judge[x][y - 1] = false;
            if(my_exist(board, word, x, y - 1, n + 1, judge))
                return true;
            judge[x][y - 1] = true;
        }

        return false;
    }
}