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
    public ListNode swapNodes(ListNode head, int k) {
        ListNode first = head;
        ListNode last = head;
        int n = getLength(head);
        int t = k;
        while (t-- > 1) first = first.next;
        t = n-k+1;
        while(t-- > 1) last = last.next;
        int temp = first.val;
        first.val = last.val;
        last.val = temp;
        return head;
    }

    private int getLength(ListNode node) {
        if (node == null) return 0;
        return 1 + getLength(node.next);
    }
}
