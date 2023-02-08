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
    ListNode headGlobal;
    public boolean isPalindrome(ListNode head) {
        headGlobal = head;
        return checkPalindrome(head);
    }

    public boolean checkPalindrome(ListNode head) {
        if (head != null) {
            if (!checkPalindrome(head.next)) {
                return false;
            }
            if (head.val != headGlobal.val) {
                return false;
            }
            headGlobal = headGlobal.next;
        }
        return true;
    }
}
