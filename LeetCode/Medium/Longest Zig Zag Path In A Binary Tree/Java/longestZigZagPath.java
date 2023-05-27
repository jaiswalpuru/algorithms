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
    int maxLen;
    public int longestZigZag(TreeNode root) {
        maxLen = 0;
        if (root.left != null) dfs(root.left, 0, 1);
        if (root.right != null) dfs(root.right, 1, 1);
        dfs(root, -1, 0);
        return maxLen;
    }

    private void dfs(TreeNode root, int dir, int currLen) {
        if (root == null) {
            return;
        }

        maxLen = Math.max(maxLen, currLen);

        if (dir == 0) {
            dfs(root.right, 1, currLen+1);
            dfs(root.left, 0, 1);
        } else {
            dfs(root.left, 0, currLen+1);
            dfs(root.right, 1, 1);
        }

    }
}
