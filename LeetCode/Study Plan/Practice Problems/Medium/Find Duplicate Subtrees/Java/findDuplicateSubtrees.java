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
    public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
    List<TreeNode> res = new ArrayList<>();
        dfs(root, new HashMap<>(), res);
        return res;
    }

    public String dfs(TreeNode node, Map<String, Integer> visited, List<TreeNode> res) {
        if (node == null) return "";

        String str = node.val + ":" + dfs(node.left, visited, res) + ":" + dfs(node.right, visited, res);
        visited.put(str, visited.getOrDefault(str, 0)+1);
        if (visited.get(str) == 2) res.add(node);

        return str; 
    }
}
