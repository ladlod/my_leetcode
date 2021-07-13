package common_list;

import java.util.Scanner;
import java.util.Stack;
public class leetcode71{

    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        String path = in.nextLine();
        System.out.println(simplifyPath(path));
        in.close();
    }
    
    static String simplifyPath(String path) {
        String[] pathS = path.split("/");
        Stack<String> s = new Stack<>();

        for(int i = 0; i < pathS.length; i++){
            if(pathS[i].equals("."))
                continue;
            else if(pathS[i].equals("..")){
                if(s.empty())
                    continue;
                else
                    s.pop();
            }
            else{
                if(pathS[i].length() != 0)
                    s.push(pathS[i]);
            }
        }

        String res = "";

        while(!s.empty()){
            res = "/" + s.pop() + res;
        }

        return res.length() == 0 ? "/" : res;
    }
    /*static String cdPath(int now, int flag, String[] pathS){
        if(flag == pathS.length)
            return now == -1? "" : pathS[now];
        else{
            if(pathS[flag].equals("."))
                return cdPath(now, flag + 1, pathS);
            else if(pathS[flag].equals(".."))
                return cdPath(now - 1, flag + 1, pathS);
            else
                return cdPath(now + 1, flag + 1, pathS);
        }
    }*/
}