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
    static int good;
    public int goodNodes(TreeNode root) {
        good = 1;
        recur(root, root.val);
        return good;
    }

    private static void recur(TreeNode root, int val) {
        if (root == null) {
            return;
        }
        if (root.left != null && val <= root.left.val) {
            good++;
            recur(root.left, root.left.val);
        } else {
            recur(root.left, val);
        }

        if (root.right != null && val <= root.right.val) {
            good++;
            recur(root.right, root.right.val);
        } else {
            recur(root.right, val);
        }
    }
}
