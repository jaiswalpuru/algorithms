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
    ListNode h;
    int sum;
    public int pairSum(ListNode head) {
        h = head;
        sum = 0;
        twinSum(head);
        return sum;
    }

    private void twinSum(ListNode head) {
        if (head == null) return;
        twinSum(head.next);
        sum = Math.max(sum, h.val+head.val);
        h = h.next;
    }
}
