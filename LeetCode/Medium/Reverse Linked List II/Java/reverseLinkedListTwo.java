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
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (head.next == null) return head;
        ListNode temp = head;
        ListNode prev = new ListNode(-1);
        prev.next = head;
        ListNode res = prev;
        ArrayDeque<ListNode> st = new ArrayDeque<>();
        int l = left;
        while (l-- > 1) {
            prev = temp;
            temp = temp.next;
        }
        right -= left;
        while(right-- >= 0) {
            st.push(temp);
            temp = temp.next;
        }
        while (!st.isEmpty()) {
            prev.next = st.pop();
            prev = prev.next;
        }
        prev.next = temp;
        return res.next;
    }
}
