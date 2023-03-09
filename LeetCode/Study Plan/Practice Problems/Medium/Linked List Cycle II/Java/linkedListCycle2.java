/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null) return null;

        ListNode intersectingNode = getIntersectingNode(head);
        if (intersectingNode == null) return null;

        ListNode p1 = head;
        ListNode p2 = intersectingNode;
        while (p1 != p2) {
            p1 = p1.next;
            p2 = p2.next;
        } 

        return p1;
    }

    public ListNode getIntersectingNode(ListNode head) {
        ListNode sp = head;
        ListNode fp = head;

        while(fp != null && fp.next != null) {
            sp = sp.next;
            fp = fp.next.next;
            if (sp == fp) return sp;
        }

        return null;
    }
}
