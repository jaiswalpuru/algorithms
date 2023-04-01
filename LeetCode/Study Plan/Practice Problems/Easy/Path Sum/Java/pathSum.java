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
    public boolean hasPathSum(TreeNode root, int targetSum) {
        if (root == null) return false;
        return inorder(root, targetSum, 0);
    }

    private boolean inorder(TreeNode root, int target, int sum) {
        if (root == null) {
            return false;
        }

        if (root.left == null && root.right == null) return root.val + sum == target;

        return inorder(root.left, target, sum+root.val) || inorder(root.right, target, sum+root.val);
    }
}
