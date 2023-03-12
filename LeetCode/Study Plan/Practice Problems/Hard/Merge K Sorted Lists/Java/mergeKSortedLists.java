/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) return null;
        
        int size = lists.length-1;

        while (size != 0) {
            int i = 0;
            int j = size;
            while(i < j) {
                lists[i] = mergeTwoList(lists[i], lists[j]);
                i++;
                j--;
                if (i >= j) size = j;
            }
        }

        return lists[0];
    }

    public ListNode mergeTwoList(ListNode a, ListNode b) {
        if (a == null) return b;
        if (b == null) return a;

        ListNode c = new ListNode(-1);
        ListNode res = c;
        while (a != null && b != null) {
            if (a.val > b.val) {
                c.next = new ListNode(b.val);
                b = b.next;
            }else {
                c.next = new ListNode(a.val);
                a = a.next;
            }
            c = c.next;
        }

        while (a != null) {
            c.next = new ListNode(a.val);
            a = a.next;
            c = c.next;
        }

        while (b != null) {
            c.next = new ListNode(b.val);
            b = b.next;
            c = c.next;
        }

        return res.next;
    }
}
