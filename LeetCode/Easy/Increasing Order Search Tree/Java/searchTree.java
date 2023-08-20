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
    List<TreeNode> in;
    public TreeNode increasingBST(TreeNode root) {
        in = new ArrayList<>();
        inorder(root);
        for (int i=1; i<in.size(); i++) {
            in.get(i-1).left = null;
            in.get(i-1).right = in.get(i);
        }
        in.get(in.size()-1).left = null;
        in.get(in.size()-1).right = null;
        return in.get(0);
    }

    private void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        in.add(root);
        inorder(root.right);
    }
}
