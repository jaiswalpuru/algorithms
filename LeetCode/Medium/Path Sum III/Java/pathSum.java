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
    int cnt;
    Map<Long, Integer> sumMap;
    public int pathSum(TreeNode root, int targetSum) {
        cnt = 0;
        sumMap = new HashMap<>();
        recur(root, targetSum, 0);
        return cnt;
    }

    private void recur(TreeNode root, int targetSum, long sum) {
        if (root == null) return;
        sum += root.val;
        if (sum == targetSum) cnt++;
        cnt += sumMap.getOrDefault(sum-targetSum, 0);
        sumMap.put(sum, sumMap.getOrDefault(sum, 0)+1);
        recur(root.left, targetSum, sum);
        recur(root.right, targetSum, sum);
        sumMap.put(sum, sumMap.get(sum)-1);
    }
}
