package common_list;

import java.util.Scanner;
//import java.util.List;;


public class leetcode19 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ListNode head = new ListNode(0);
        ListNode p = head;
        int len = in.nextInt();
        while (len-- > 0) {
            int x = in.nextInt();
            ListNode tmp = new ListNode(x);
            p.next = tmp;
            p = p.next;
        }
        int n = in.nextInt();
        head = removeNthFromEnd(head.next, n);
        while (head != null) {
            System.out.println(head.val);
            head = head.next;
        }
        in.close();
    }

    static ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode tmp = new ListNode(0);
        tmp.next = head;
        ListNode end = tmp, start = tmp;
        for (int i = 0; i <= n; i++) {
            end = end.next;
        }
        while (end != null) {
            start = start.next;
            end = end.next;
        }
        start.next = start.next.next;
        return tmp.next;
    }
}