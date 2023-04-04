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
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        List<List<Integer>> res = new ArrayList<>();
        recur(root, targetSum, res, new ArrayList<>(), 0);
        return res;
    }

    private void recur(TreeNode root, int target, List<List<Integer>> res, List<Integer> temp, int currSum) {
        if (root == null) return;

        if (root.left == null && root.right == null) {
            if (currSum+root.val == target) {
                temp.add(root.val);
                res.add(new ArrayList<>(temp));
                temp.remove(temp.size()-1);
            }
            return;
        }

        temp.add(root.val);
        recur(root.left, target, res, temp, currSum+root.val);
        recur(root.right, target, res, temp, currSum+root.val);
        temp.remove(temp.size()-1);
    }
}
