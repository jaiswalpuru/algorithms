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
    public boolean isValidBST(TreeNode root) {
        return recur(root, Double.NEGATIVE_INFINITY, Double.POSITIVE_INFINITY);
    }

    private boolean recur(TreeNode root, Double lowerBound, Double upperBound) {
        if (root == null) return true;
        return (root.val > lowerBound) &&
                (root.val < upperBound) &&
                recur(root.left, lowerBound, (double)root.val) &&
                recur(root.right, (double)root.val, upperBound);
    }
}
