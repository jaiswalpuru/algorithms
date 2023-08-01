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
    int good;
    public int goodNodes(TreeNode root) {
        good = 1;
        _goodNodes(root, root.val);
        return good;
    }

    private void _goodNodes(TreeNode root, int val) {
        if (root == null) return;
        if (root.left != null && root.left.val >= val) {
            good++;
            _goodNodes(root.left, root.left.val);
        } else {
            _goodNodes(root.left, val);
        }

        if (root.right != null && root.right.val >= val) {
            good++;
            _goodNodes(root.right, root.right.val);
        } else {
            _goodNodes(root.right, val);
        }
    }
}
