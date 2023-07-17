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
        ArrayDeque<ListNode> s1 = new ArrayDeque<>();
        ArrayDeque<ListNode> s2 = new ArrayDeque<>();
        
        while(l1 != null) {
            s1.push(l1);
            l1 = l1.next;
        }
        
        while(l2 != null) {
            s2.push(l2);
            l2 = l2.next;
        }

        int carry = 0;
        ListNode res = null;
        while(!s1.isEmpty() && !s2.isEmpty()) {
            int sum = s1.pop().val + s2.pop().val + carry;
            carry = sum/10;
            if (res == null) {
                res = new ListNode(sum%10);
            } else {
                ListNode t = new ListNode(sum%10);
                t.next = res;
                res = t;
            }
        }

        while(!s1.isEmpty()) {
            int sum = s1.pop().val + carry;
            carry = sum/10;
            ListNode t = new ListNode(sum%10);
            t.next = res;
            res = t;
        }

        while(!s2.isEmpty()) {
            int sum = s2.pop().val + carry;
            carry = sum/10;
            ListNode t = new ListNode(sum%10);
            t.next = res;
            res = t;
        }

        while(carry > 0) {
            ListNode t = new ListNode(carry%10);
            carry /= 10;
            t.next = res;
            res = t;
        }

        return res;
    }
}
