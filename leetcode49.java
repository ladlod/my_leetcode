import java.util.HashMap;
import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class leetcode49 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        in.nextLine();
        String[] strs = new String[n];
        for (int i = 0; i < n; i++) {
            strs[i] = in.nextLine();
        }
        System.out.println(groupAnagrams(strs));
        in.close();
    }

    static List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String>> map = new HashMap<>();
        for (int i = 0; i < strs.length; i++) {
            char[] ctmp = strs[i].toCharArray();
            Arrays.sort(ctmp);
            String stmp = String.copyValueOf(ctmp);
            if (map.get(stmp) == null) {
                List<String> ltmp = new ArrayList<>();
                ltmp.add(strs[i]);
                map.put(stmp, ltmp);
            } else {
                List<String> ltmp = map.get(stmp);
                ltmp.add(strs[i]);
            }
        }

        return new ArrayList<List<String>>(map.values());
    }
}