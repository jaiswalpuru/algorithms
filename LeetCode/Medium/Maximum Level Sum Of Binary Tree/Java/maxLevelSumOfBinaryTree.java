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
        ArrayDeque<TreeNode> q = new ArrayDeque<>();
        q.offer(root);
        int lvl = 0;
        int l = 0;
        int sum = Integer.MIN_VALUE;
        while (!q.isEmpty()) {
            l++;
            int size = q.size();
            int lvlSum = 0;
            for (int i=0; i<size; i++) {
                TreeNode curr = q.poll();
                lvlSum += curr.val;
                if (curr.left != null) q.offer(curr.left);
                if (curr.right != null) q.offer(curr.right);
            }
            if (sum < lvlSum) {
                sum = lvlSum;
                lvl = l;
            }
        }
        return lvl;
    }
}
