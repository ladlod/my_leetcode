import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;

public class leetcode30 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        int n = in.nextInt();
        in.nextLine();
        String[] words = new String[n];
        int i = 0;
        while (i < n) {
            words[i] = in.nextLine();
            i++;
        }
        System.out.println(findSubstring(s, words));
        in.close();
    }

    static List<Integer> findSubstring(String s, String[] words) {
        List<Integer> res = new ArrayList<>();
        if (s == null || s.length() == 0 || words.length == 0 || words == null)
            return res;
        int wlen = words[0].length();
        int wnum = words.length;
        HashMap<String, Integer> wordmap = new HashMap<>();
        for (int i = 0; i < wnum; i++) {
            if (wordmap.get(words[i]) == null)
                wordmap.put(words[i], 1);
            else {
                int tmp = wordmap.get(words[i]);
                wordmap.remove(words[i]);
                wordmap.put(words[i], tmp + 1);
            }
        }

        HashMap<Integer, Integer> map = new HashMap<>();
        for (int j = 0; j <= s.length() - wlen; j++) {
            for (int i = 0; i < words.length; i++) {
                if (s.substring(j, j + wlen).equals(words[i])) {
                    map.put(j, i);
                    break;
                }
            }
        }
        System.out.println("map = " + map);
        System.out.println("wordmap = " + wordmap);
        for (int i = 0; i < s.length(); i++) {
            HashMap<String, Integer> tmpmap = (HashMap) wordmap.clone();
            int p = i;
            int count = 0;
            while (map.get(p) != null && count < words.length) {
                int x = map.get(p);
                if (tmpmap.get(words[x]) != null) {
                    int tmp = (int) tmpmap.get(words[x]);
                    tmpmap.remove(words[x]);
                    if (tmp > 1)
                        tmpmap.put(words[x], tmp - 1);
                }
                p += wlen;
                count++;
            }
            if (tmpmap.isEmpty())
                res.add(i);
        }
        return res;
    }
}