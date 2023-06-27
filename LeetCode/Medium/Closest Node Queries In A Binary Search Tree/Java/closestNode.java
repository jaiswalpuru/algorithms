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
    public List<List<Integer>> closestNodes(TreeNode root, List<Integer> queries) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> sorted = new ArrayList<>();
        inorder(root, sorted);
        for (Integer q : queries) {
            List<Integer> indexes = new ArrayList<>();
            int ind = Collections.binarySearch(sorted, q);
            if (ind >= 0) {
                indexes.add(sorted.get(ind));
                indexes.add(sorted.get(ind));
                res.add(indexes);
            } else {
                ind = Math.abs(ind)-2;
                if (ind < 0) indexes.add(-1);
                else indexes.add(sorted.get(ind));
                ind = ind+1;
                if (ind >= sorted.size()) indexes.add(-1);
                else indexes.add(sorted.get(ind));
                res.add(indexes);
            }
        }
        return res;
    }

    private void inorder(TreeNode root, List<Integer> sorted) {
        if (root == null) return;
        inorder(root.left, sorted);
        sorted.add(root.val);
        inorder(root.right, sorted);
    }
}
