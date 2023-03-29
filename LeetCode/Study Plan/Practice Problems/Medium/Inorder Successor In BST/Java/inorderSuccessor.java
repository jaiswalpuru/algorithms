/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    List<TreeNode> order;
    public TreeNode inorderSuccessor(TreeNode root, TreeNode p) {        
        order = new ArrayList<>();
        inorder(root);
        for (int i=0; i<order.size(); i++) {
            if (order.get(i).val == p.val) {
                if (i+1 < order.size()) return order.get(i+1);
            }
        }
        return null;
    }

    public void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        order.add(root);
        inorder(root.right);
    }

}
