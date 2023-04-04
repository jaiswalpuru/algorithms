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
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> rightSide = new ArrayList<>();
        if (root == null) return rightSide;
        ArrayDeque<TreeNode> q = new ArrayDeque<>();
        q.offer(root);

        while(!q.isEmpty()) {
            int qSize = q.size();
            rightSide.add(q.getLast().val);
            for (int i=0; i<qSize; i++) {
                TreeNode curr = q.poll();
                if (curr.left != null) q.offer(curr.left);
                if (curr.right != null) q.offer(curr.right);
            }
        }

        return rightSide;
    }
}
