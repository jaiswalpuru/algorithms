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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode res = new ListNode();
        ListNode ans = res;
        int carry = 0;

        while(l1 != null && l2 != null) {
            int sum = l1.val + l2.val + carry;
            res.next = new ListNode(sum%10);
            res = res.next;
            l1 = l1.next;
            l2 = l2.next;
            carry = sum/10;
        }

        while(l1 != null) {
            int sum = carry + l1.val;
            res.next = new ListNode(sum%10);
            res = res.next;
            carry = sum/10;
            l1 = l1.next;
        }

        while(l2 != null) {
            int sum = carry + l2.val;
            res.next = new ListNode(sum%10);
            res = res.next;
            carry = sum/10;
            l2 = l2.next;
        }

        while(carry > 0) {
            res.next = new ListNode(carry%10);
            carry = carry/10;
            res = res.next;
        }

        return ans.next;
    }
}
