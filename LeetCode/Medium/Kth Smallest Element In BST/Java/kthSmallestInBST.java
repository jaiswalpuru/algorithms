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
    List<Integer> in;
    public int kthSmallest(TreeNode root, int k) {
        in = new ArrayList<>();
        inorder(root);
        return in.get(k-1);
    }

    private void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        in.add(root.val);
        inorder(root.right);
    }
}
