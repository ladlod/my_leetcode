package common_list;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class leetcode17 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String digits = in.nextLine();
        System.out.println(letterCombinations(digits));
        in.close();
    }

    static List<String> letterCombinations(String digits) {
        List<String> letter = new ArrayList<>();
        if (digits.length() == 0)
            return letter;
        HashMap<String, String> map = new HashMap<>();
        map.put("2", "abc");
        map.put("3", "def");
        map.put("4", "ghi");
        map.put("5", "jkl");
        map.put("6", "mno");
        map.put("7", "pqrs");
        map.put("8", "tuv");
        map.put("9", "wxyz");
        for (int j = 0; j < map.get(digits.substring(0, 1)).length(); j++) {
            letter.add("" + map.get(digits.substring(0, 1)).charAt(j));
        }
        for (int i = 1; i < digits.length(); i++) {
            String words = map.get(digits.substring(i, i + 1));
            int len = letter.size();
            for (int j = 0; j < len; j++) {
                System.out.println(letter);
                String tmp = letter.remove(0);
                // System.out.println(tmp);
                for (int k = 0; k < words.length(); k++)
                    letter.add(tmp + words.charAt(k));
            }
        }
        // System.out.println(letter);
        return letter;
    }
}