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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode fp = head;
        while(n-- > 0) {
            fp = fp.next;
        }

        if (fp == null) return head.next;
        ListNode sp = head;
        ListNode prev = head;
        while (fp != null) {
            prev = sp;
            sp = sp.next;
            fp = fp.next;
        }
        prev.next = sp.next;
        return head;
    }
}
