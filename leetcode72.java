import java.util.Scanner;
public class leetcode72{
    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        String word1 = in.nextLine(), word2 = in.nextLine();
        System.out.println(minDistance(word1, word2));
        in.close();
    }

    static int minDistance(String word1, String word2){
        int len1 = word1.length();
        int len2 = word2.length();
        int[][] minD = new int[len1 + 1][len2 + 1];
        minD[0][0] = 0;
        for(int i = 1; i <= len1; i++)
            minD[i][0] = i;
        for(int i = 1; i <= len2; i++)
            minD[0][i] = i;
        for(int i = 1; i <= len1; i++){
            for(int j = 1; j <= len2; j++){
                if(word1.charAt(i - 1) == word2.charAt(j - 1))
                    minD[i][j] = minD[i - 1][j - 1];
                else{
                    minD[i][j] = minOf3(minD[i - 1][j], minD[i][j - 1], minD[i - 1][j - 1]) + 1;
                }
            }
        }
        /*for(int i = 0; i <= len1; i++){
            for(int j = 0; j <= len2; j++){
                System.out.print(minD[i][j] + " ");
            }
            System.out.println();
        }*/
        
        return minD[len1][len2];
    }

    static int minOf3(int num1, int num2, int num3){
        int tmp = num1 < num2 ? num1 : num2;
        return tmp < num3 ? tmp : num3; 
    }
}