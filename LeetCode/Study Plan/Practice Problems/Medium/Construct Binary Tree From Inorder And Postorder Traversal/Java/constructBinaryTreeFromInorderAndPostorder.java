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
    int[] inorder;
    int[] postorder;
    int postInd=0;
    HashMap<Integer, Integer> inorderIndex = new HashMap<>();
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        this.inorder = inorder;
        this.postorder = postorder;
        
        for (int i=0; i<inorder.length; i++) {
            inorderIndex.put(inorder[i], i);
        }

        postInd = postorder.length-1;
        return recur(0, inorder.length-1);
    }

    public TreeNode recur(int left, int right) {
        if (left > right) return null;
        
        TreeNode root = new TreeNode(this.postorder[this.postInd]);
        int ind = inorderIndex.get(root.val);
        this.postInd--;
        
        root.right = recur(ind+1, right);
        root.left = recur(left, ind-1);
        return root;
    }
}
