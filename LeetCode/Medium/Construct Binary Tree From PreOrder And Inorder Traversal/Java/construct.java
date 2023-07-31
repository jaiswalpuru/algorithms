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
    Map<Integer, Integer> indexIn;
    int preIndex;
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        preIndex = 0;
        indexIn = new HashMap<>();
        for (int i=0; i<inorder.length; i++) indexIn.put(inorder[i], i);
        return preRecur(preorder, 0, preorder.length-1);
    }

    private TreeNode preRecur(int[] preorder, int left, int right) {
        if (left > right) return null;
        int rootVal = preorder[preIndex++];
        TreeNode root = new TreeNode(rootVal);
        root.left = preRecur(preorder, left, indexIn.get(rootVal)-1);
        root.right = preRecur(preorder, indexIn.get(rootVal)+1, right);
        return root;
    }
}
