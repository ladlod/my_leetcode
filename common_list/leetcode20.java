package common_list;

import java.util.Scanner;
import java.util.Stack;

public class leetcode20 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String str = in.nextLine();
        System.out.println(isValid(str));
        in.close();
    }

    static boolean isValid(String s) {
        Stack<Character> S = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(' || c == '[' || c == '{') {
                S.push(c);
                System.out.println(S);
            } else if (c == ')') {
                if (S.isEmpty() || S.pop() != '(')
                    return false;
                System.out.println(S);
            } else if (c == ']') {
                if (S.isEmpty() || S.pop() != '[')
                    return false;
                System.out.println(S);
            } else if (c == '}') {
                if (S.isEmpty() || S.pop() != '{')
                    return false;
                System.out.println(S);
            }
        }
        if (S.isEmpty())
            return true;
        else
            return false;
    }
}