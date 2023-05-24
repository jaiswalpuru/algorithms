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
    public TreeNode deleteNode(TreeNode root, int key) {
        TreeNode temp = root;
        TreeNode prev = new TreeNode(-1, root, null);
        TreeNode res = prev;
        // this will break after it finds the node with key
        while(temp != null) {
            if (temp.val > key) {
                prev = temp;
                temp = temp.left;
            } else if (temp.val < key) {
                prev = temp;
                temp = temp.right;
            } else {
                break;
            }
        }
        if (temp == null) return root; // key not found
        if (temp.left == null && temp.right == null) { // if leaf
            if (prev.left == temp) prev.left = null;
            else prev.right = null;
        } else if (temp.left == null) {
            System.out.println(temp.val);
            if (prev.left == temp) prev.left = temp.right;
            else prev.right = temp.right;
        } else if (temp.right == null) {
            if (prev.left == temp) prev.left = temp.left;
            else prev.right = temp.left;
        } else {

            TreeNode succ = getSuccessor(temp);
            int t = temp.val;
            temp.val = succ.val;
            succ.val = t;
            prev = temp;
            temp = temp.right;
            while (temp.left != null) {
                prev = temp;
                temp = temp.left;
            }
            if (prev.left == temp) prev.left = temp.right;
            else prev.right = temp.right;
        }

        return res.left;
    }

    // to get the successor of the current node
    private TreeNode getSuccessor(TreeNode node) {
        node = node.right;
        while(node.left != null) node = node.left;
        return node;
    }
}
