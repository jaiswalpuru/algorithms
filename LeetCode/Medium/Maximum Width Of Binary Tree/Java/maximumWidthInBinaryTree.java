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
    public int widthOfBinaryTree(TreeNode root) {
        if (root == null) return 0;
        LinkedList<Pair<TreeNode, Integer>> q = new LinkedList<>();
        Integer maxWidth = 0;

        q.addLast(new Pair<>(root, 0));
        while(q.size() > 0) {
            Pair<TreeNode, Integer> front = q.getFirst();
            
            int currLevelSize = q.size();
            Pair<TreeNode, Integer> temp = null;
            for (int i=0; i<currLevelSize; i++) {
                temp = q.removeFirst();
                TreeNode node = temp.getKey();
                if (node.left != null) q.addLast(new Pair<>(node.left, 2*temp.getValue()));
                if (node.right != null) q.addLast(new Pair<>(node.right, 2*temp.getValue()+1));
            }
            maxWidth = Math.max(maxWidth, temp.getValue()-front.getValue()+1);
        }
        return maxWidth;
    }
}
