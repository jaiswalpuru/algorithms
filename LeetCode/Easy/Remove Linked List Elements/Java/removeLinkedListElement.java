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
    public ListNode removeElements(ListNode head, int val) {
        if (head == null) return head;
        ListNode tempHead = head;
        ListNode prev = null;
        while(head != null) {
            if (head.val == val) {
                if (prev == null) {
                    tempHead = head.next;
                    head = head.next;
                } else {
                    prev.next = head.next;
                    head = head.next;
                }
            } else {
                prev = head;
                head = head.next;
            }
        }
        return tempHead;
    }
}
