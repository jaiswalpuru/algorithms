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
    public ListNode removeNodes(ListNode head) {
        if (head == null) return head;
        ArrayDeque<ListNode> stack = new ArrayDeque<>();
        while(head != null) {
            if (stack.isEmpty()) {
                stack.push(head);
                head = head.next;
            } else {
                while(!stack.isEmpty() && stack.peek().val < head.val) stack.pop();
                stack.push(head);
                head = head.next;
            }
        }

        ListNode res = new ListNode(-1);
        ListNode temp = stack.removeLast();
        res.next = temp;
        while(!stack.isEmpty()) {
            temp.next = stack.removeLast();
            temp = temp.next;
        }
        temp.next = null;
        return res.next;
    }
}

// recursive solution
class Solution {
    public ListNode removeNodes(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode p = removeNodes(head.next);
        head.next = p;
        if (head.val < head.next.val) return head.next;
        return head;
    }
}
