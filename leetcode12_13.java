import java.util.HashMap;
import java.util.Scanner;

public class leetcode12_13 {
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int num = in.nextInt();
        System.out.println(intToRoman(num));
        in.nextLine();
        String roman = in.nextLine();
        System.out.println(romanToInt(roman));
        in.close();
    }

    public static String intToRoman(int num) {
        String str = "";
        int[] intnum = { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 };
        String[] roman = { "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" };
        for (int i = 0; i < intnum.length; i++) {
            int n = num / intnum[i];
            for (int j = 1; j <= n; j++) {
                str += roman[i];
            }
            num = num - n * intnum[i];
        }
        return str;
    }

    public static int romanToInt(String s) {
        int ret = 0;
        HashMap<String, Integer> map = new HashMap<>();
        map.put("M", 1000);
        map.put("CM", 900);
        map.put("D", 500);
        map.put("CD", 400);
        map.put("C", 100);
        map.put("XC", 90);
        map.put("L", 50);
        map.put("XL", 40);
        map.put("X", 10);
        map.put("IX", 9);
        map.put("V", 5);
        map.put("IV", 4);
        map.put("I", 1);
        for (int i = 0; i < s.length(); i++) {
            if (i != s.length() - 1 && map.containsKey(s.substring(i, i + 2))) {
                ret += map.get(s.substring(i, i + 2));
                i++;
            } else {
                ret += map.get(s.substring(i, i + 1));
            }
        }
        return ret;
    }
}