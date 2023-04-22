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
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode sortedListToBST(ListNode head) {
        List<Integer> sortedList = new ArrayList<>();

        while(head != null) {
            sortedList.add(head.val);
            head = head.next;
        }

        return bst(sortedList, 0, sortedList.size()-1);
    }

    public TreeNode bst(List<Integer> arr, int l, int r) {
        if (l > r) return null;

        int mid = l + (r-l)/2;
        TreeNode root = new TreeNode(arr.get(mid));
        root.left = bst(arr, l, mid-1);
        root.right = bst(arr, mid+1, r);
        return root;
    }
}
