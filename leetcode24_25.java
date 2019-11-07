import java.util.Scanner;
//import java.util.List;;

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}

public class leetcode24_25 {
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
        /*
         * head = swapPairs(head); while (head != null) { System.out.println(head.val);
         * head = head.next; }
         */
        int k = in.nextInt();
        head = reserseKGroup(head.next, k);
        while (head != null) {
            System.out.println(head.val);
            head = head.next;
        }
        in.close();
    }

    static ListNode swapPairs(ListNode head) {
        ListNode p = new ListNode(0);
        p.next = head;
        ListNode ret = p;
        while (p.next != null && p.next.next != null) {
            ListNode tmp1 = p.next;
            ListNode tmp2 = p.next.next;
            p.next = tmp2;
            tmp1.next = tmp2.next;
            tmp2.next = tmp1;
            p = tmp1;
        }
        return ret.next;
    }

    static ListNode reserseKGroup(ListNode head, int k) {
        ListNode p = new ListNode(0);
        p.next = head;
        ListNode ret = p;
        w: while (p.next != null) {
            // System.out.println(p.val);
            ListNode start = p;
            ListNode end = p;
            for (int i = 0; i <= k; i++) {
                if (end == null)
                    break w;
                end = end.next;
            }
            p = reverseGroup(p.next, k);
            start.next = p;
            for (int i = 1; i < k; i++) {
                p = p.next;
            }
            p.next = end;
        }
        return ret.next;
    }

    static ListNode reverseGroup(ListNode head, int n) {
        if (head.next == null || n == 1)
            return head;
        n--;
        ListNode ret = reverseGroup(head.next, n);
        head.next.next = head;
        head.next = null;
        return ret;
    }
}