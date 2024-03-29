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
    int val;
    public int minDepth(TreeNode root) {
        if (root == null) return 0;
        return recur(root);
    }

    private int recur(TreeNode root) {
        if (root == null) return Integer.MAX_VALUE;
        if (root.left == null && root.right == null) return 1;
        int left = recur(root.left);
        int right = recur(root.right);
        return 1+Math.min(left, right);
    }
}
