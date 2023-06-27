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
    int minVal;
    double diff;
    public int closestValue(TreeNode root, double target) {
        minVal = -1;
        diff = Double.MAX_VALUE;
        recur(root, target);
        return minVal;
    }

    private void recur(TreeNode root, double target) {
        if (root == null) return;
        if (diff >= Math.abs(target-root.val)) {
            if (diff == Math.abs(target-root.val)) {
                if (root.val < minVal) minVal = root.val;
            } else {
                diff = Math.abs(target-root.val);
                minVal = root.val;
            }
        }
        recur(root.left, target);
        recur(root.right, target);
    }
}
