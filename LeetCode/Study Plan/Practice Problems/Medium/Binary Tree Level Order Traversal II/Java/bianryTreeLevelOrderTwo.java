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
    List<List<Integer>> order;
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        order = new ArrayList<>();
        
        if (root == null) return order;
        
        Queue<TreeNode> q = new ArrayDeque<>();
        q.offer(root);

        while(!q.isEmpty()) {
            int qSize = q.size();
            List<Integer> currLevel = new ArrayList<>();
            for(int i=0; i<qSize; i++) {
                TreeNode curr = q.poll();
                currLevel.add(curr.val);
                if (curr.left != null) q.offer(curr.left);
                if (curr.right != null) q.offer(curr.right);
            }
            order.add(currLevel);
        }
        Collections.reverse(order);
        return order;
    }
}
