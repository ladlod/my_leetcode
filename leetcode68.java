import java.util.List;
import java.util.ArrayList;
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
        System.out.println(dealWords(words, maxWidth));
        in.close();
    }
    static List<String> dealWords(String[] words, int maxWidth){
        List<String> res = new ArrayList<>();
        int i = 0, l = 0, flag = 0;
        while(i < words.length){
            if(l + words[i].length() <= maxWidth){
                l = l + words[i].length() + 1;
                //System.out.println("l=" + l);
            }
            else{
                //System.out.println("start=" + flag + " end=" + i);
                res.add(MaxW(words, flag, i, maxWidth));
                flag = i;
                l = words[i].length() + 1;
            }
            i++;
        }
        String last = "";
        int len = 0;
        for(int j = flag; j < words.length; j++){
            last = last + words[j] + " ";
            len = len + words[j].length() + 1;
        }
        for(int j = len - 1; j < maxWidth; j++)
                last += " ";
        res.add(last);

        return res;       
    }

    static String MaxW(String[] words, int start, int end, int maxWidth){
        String res = "";
        if(end - start > 1){
            int len = 0;
            for(int i = start; i < end; i++)
                len += words[i].length();
            int insert_num = (maxWidth - len) / (end - start - 1);
            int add_num = (maxWidth - len) % (end - start - 1);

            for(int i  = start; i < end; i++){
                res += words[i];
                //System.out.print(words[i]);
                int j = 0;
                if(i != end - 1){
                    while(j < insert_num + (add_num > 0 ? 1 : 0)){
                        res += " ";
                        j++;
                    }
                    add_num--;
                }
            }
        }
        else{
            res += words[start];
            //System.out.print(words[start]);
            for(int i = words[start].length(); i < maxWidth; i++)
                res += " ";
        }
        return res;
    }
}