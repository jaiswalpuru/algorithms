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
    public ListNode deleteMiddle(ListNode head) {
        ListNode sp = head;
        ListNode fp = head;
        ListNode prev = new ListNode(-1); //sentinel node
        prev.next = head;
        while (fp != null && fp.next != null) {
            prev = sp;
            sp = sp.next;
            fp = fp.next.next;
        }
        prev.next = sp.next;
        return prev.val == -1 ? null : head;
    }
}
