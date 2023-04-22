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
    int sum = 0;
    public boolean checkTree(TreeNode root) {
        recur(root);
        return root.val == sum;
    }

    public void recur(TreeNode root) {
        if (root == null) return;
        if (root.left == null && root.right == null) {
            sum += root.val;
        }

        recur(root.left);
        recur(root.right);
    }
}
