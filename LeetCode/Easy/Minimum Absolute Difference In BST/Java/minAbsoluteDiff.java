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
    List<Integer> order;
    public int getMinimumDifference(TreeNode root) {
        order = new ArrayList<>();
        inOrder(root);
        int diff = (int)1e9+7;
        for (int i=1; i<order.size(); i++)
            diff = Math.min(diff, Math.abs(order.get(i)-order.get(i-1)));

        return diff;
    }

    private void inOrder(TreeNode root) {
        if (root == null) return;
        inOrder(root.left);
        order.add(root.val);
        inOrder(root.right);
    }
}
