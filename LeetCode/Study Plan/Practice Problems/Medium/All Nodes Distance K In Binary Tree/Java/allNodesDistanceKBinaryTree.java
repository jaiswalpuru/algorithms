/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    List<Integer> res;

    public List<Integer> distanceK(TreeNode root, TreeNode target, int k) {
        res = new ArrayList<>();
        dfs(root, target, k);
        return res;
    }

    private int dfs(TreeNode root, TreeNode target, int dis) {
        if (root == null) return -1;
        if (root == target) {
            subTreeAdd(root, dis, 0);
            return 1;
        } else {
            int left = dfs(root.left, target, dis);
            int right = dfs(root.right, target, dis);

            if (left != -1) {
                if (left == dis) res.add(root.val);
                subTreeAdd(root.right, dis, left+1);
                return left+1;
            } else if (right != -1) {
                if (right == dis) res.add(root.val);
                subTreeAdd(root.left, dis, right+1);
                return right+1;
            } else {
                return -1;
            }
        }
    }

    private void subTreeAdd(TreeNode root, int dis, int currDis) {
        if (root == null) return;
        if (dis == currDis) {
            res.add(root.val);
        }
        subTreeAdd(root.left, dis, currDis+1);
        subTreeAdd(root.right, dis, currDis+1);
    }
}
