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
    List<Integer> sortedElement;

    public int minDiffInBST(TreeNode root) {
        sortedElement = new ArrayList<>();
        inorder(root);

        int minDiff = sortedElement.get(1) - sortedElement.get(0);
        for (int i=2; i<sortedElement.size(); i++) 
            minDiff = Math.min(minDiff, sortedElement.get(i)-sortedElement.get(i-1));
        
        return minDiff;
    }

    public void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        sortedElement.add(root.val);
        inorder(root.right);
    }
}
