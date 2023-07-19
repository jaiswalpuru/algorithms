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
    public ListNode mergeInBetween(ListNode list1, int a, int b, ListNode list2) {
        int index = 0;
        ListNode head = list1;
        if (index == a) {
            while(index++ != b) list1 = list1.next;
            head = list2;
            while(list2.next != null) list2 = list2.next;
            list2.next = list1.next;
        } else {
            while(index++ != a-1) list1 = list1.next;
            ListNode temp = list1.next;
            while(index++ != b) temp = temp.next;
            list1.next = list2;
            while(list2.next != null) list2 = list2.next;
            list2.next = temp.next;
        }
        return head;
    }
}
