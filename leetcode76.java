import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class leetcode76 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String S = in.nextLine();
        String T = in.nextLine();
        System.out.println(minWindow(S, T));
        in.close();
    }

    static String minWindow(String s, String t) {

        HashMap<Character, Integer> mapt = new HashMap<>();

        for(int i = 0; i < t.length(); i++){
            int count = mapt.getOrDefault(t.charAt(i), 0);
            mapt.remove(t.charAt(i));
            mapt.put(t.charAt(i), count + 1);
        }
        //System.out.println(mapt);

        HashMap<Character, Integer> mapres = new HashMap<>();
        String res = "";
        int begin = 0, end = 0;
        while(true){
            if(sHasT(mapres, mapt)){
                if(res.length() == 0 || end - begin < res.length()){
                    res = s.substring(begin, end);
                }
                int count0 = mapres.get(s.charAt(begin));
                mapres.remove(s.charAt(begin));
                if(count0 > 1)
                    mapres.put(s.charAt(begin), count0 - 1);
                begin++;
            }
            else if(end < s.length()){
                int count = mapres.getOrDefault(s.charAt(end), 0);
                mapres.put(s.charAt(end), count + 1);
                end++;
            }
            else{
                break;
            }

            //System.out.println(s.substring(begin, end) + " " + res);
        }

        return res;
    }

    static boolean sHasT(HashMap<Character, Integer> mapres, HashMap<Character, Integer>  mapt){
        for(Map.Entry<Character, Integer> entry: mapt.entrySet()){
            if(mapres.getOrDefault(entry.getKey(), 0) < entry.getValue()){
                return false;
            }
        }

        return true;
    }
}