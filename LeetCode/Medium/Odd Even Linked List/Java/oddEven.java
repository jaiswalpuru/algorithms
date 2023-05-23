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
    public ListNode oddEvenList(ListNode head) {
        ListNode tail = head;
        int n = 0, i = 1;
        if (head == null) return head;
        while(tail.next != null) {
            tail = tail.next;
            n++;
        }

        ListNode tempHead = head;
        ListNode tempTail = tail;
        ListNode prev = head;
        while(n >= 0) {
            if (i%2 == 0) {
                //even
                tempTail.next = prev.next;
                prev.next = tempHead.next;
                tempHead = tempHead.next;
                tempTail = tempTail.next;
                tempTail.next = null;
            }else {
                //odd
                prev = tempHead;
                tempHead = tempHead.next;
            }
            n--;
            i++;
        }
        return head;
    }
}
