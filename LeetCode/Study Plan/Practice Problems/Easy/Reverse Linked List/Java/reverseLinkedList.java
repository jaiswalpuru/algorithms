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
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode rev = reverse(head);
        head.next = null;
        return rev;
    }

    private ListNode reverse(ListNode head) {
        if (head == null) return head;
        ListNode temp = reverse(head.next);
        if (temp == null) return head;
        head.next.next = head;
        return temp;
    }
}
