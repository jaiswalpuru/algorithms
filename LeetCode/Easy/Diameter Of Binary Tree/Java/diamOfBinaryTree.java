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
    int d = 0;
    public int diameterOfBinaryTree(TreeNode root) {
        diam(root);
        return d;
    }

    private int diam(TreeNode root) {
        if (root == null) return 0;

        int left = diam(root.left);
        int right = diam(root.right);
        d = Math.max(d, left+right);
        return 1+Math.max(left, right);
    }
}
