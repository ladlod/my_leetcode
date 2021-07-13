package common_list;

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}

public class leetcode23 {
    public static void main(String[] args) {

    }

    static ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0)
            return null;
        else if (lists.length == 1)
            return lists[0];
        else if (lists.length == 2)
            return merge2Lists(lists[0], lists[1]);
        else {
            ListNode[] begin = new ListNode[lists.length / 2];
            for (int i = 0; i < lists.length / 2; i++) {
                begin[i] = lists[i];
            }
            ListNode[] end = new ListNode[lists.length - lists.length / 2];
            for (int i = 0; i < lists.length - lists.length / 2; i++) {
                end[i] = lists[i + lists.length / 2];
            }
            return merge2Lists(mergeKLists(begin), mergeKLists(end));
        }
    }

    static ListNode merge2Lists(ListNode l1, ListNode l2) {
        ListNode l = new ListNode(0);
        ListNode p = l;
        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                l.next = l1;
                l = l.next;
                l1 = l1.next;
            } else if (l1.val > l2.val) {
                l.next = l2;
                l = l.next;
                l2 = l2.next;
            }
        }
        if (l1 == null) {
            l.next = l2;
        } else {
            l.next = l1;
        }
        return p.next;
    }
}