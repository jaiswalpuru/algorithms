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
    public int deepestLeavesSum(TreeNode root) {
        ArrayDeque<TreeNode> dq = new ArrayDeque<>();
        dq.offer(root);
        int res = 0;
        while(!dq.isEmpty()) {
            int size = dq.size();
            int sum = 0;
            for (int i=0; i<size; i++) {
                TreeNode curr = dq.poll();
                if (curr.left != null) dq.offer(curr.left);
                if (curr.right != null) dq.offer(curr.right);
                sum += curr.val;
            }
            res = sum;
        }
        return res;
    }
}
