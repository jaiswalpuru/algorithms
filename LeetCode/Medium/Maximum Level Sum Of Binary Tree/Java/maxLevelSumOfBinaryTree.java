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
    public int maxLevelSum(TreeNode root) {
        ArrayDeque<TreeNode> dq = new ArrayDeque<>();
        dq.offer(root);
        int maxSum = root.val;
        int ans = 1;
        int lvl = 1;
        while (!dq.isEmpty()) {
            int size = dq.size();
            int lvlSum = 0;
            for (int i=0; i<size; i++) {
                TreeNode curr = dq.poll();
                lvlSum += curr.val;
                if (curr.left != null) dq.offer(curr.left);
                if (curr.right != null) dq.offer(curr.right);
            }
            if (maxSum < lvlSum) {
                maxSum = lvlSum;
                ans = lvl;
            }
            lvl++;
        }
        return ans;
    }
}
