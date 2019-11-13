import java.util.Scanner;
public class leetcode68{
    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        in.nextLine();
        String[] words = new String[n];
        for(int i = 0; i < n; i++){
            words[i] = in.nextLine();
        }
        int maxWidth = in.nextInt();
        /*for(int i = 0; i < n; i++)
            System.out.print(i + "=" + words[i] + " ");
        System.out.println();*/
        dealWords(words, maxWidth);
        in.close();
    }
    static void dealWords(String[] words, int maxWidth){
        int i = 0, l = 0, flag = 0;
        while(i < words.length){
            if(l + words[i].length() + 1 <= maxWidth)
                l = l + words[i].length() + 1;
            else{
                printMaxW(words, flag, i, maxWidth);
                flag = i;
                l = 0;
            }
            i++;
        }
        printMaxW(words, flag, words.length, maxWidth);        
    }

    static void printMaxW(String[] words, int start, int i, int maxWidth){

    }
}