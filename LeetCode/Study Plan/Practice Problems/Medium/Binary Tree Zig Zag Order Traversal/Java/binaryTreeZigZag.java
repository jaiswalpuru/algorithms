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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;

        ArrayDeque<TreeNode> queue = new ArrayDeque<>();
        queue.offer(root);
        boolean flag = false;

        while(!queue.isEmpty()) {
            int size = queue.size();
            ArrayDeque<TreeNode> currQueue = new ArrayDeque<>();
            List<Integer> currRes = new ArrayList<>();

            for (int i=0; i<size; i++) {
                TreeNode curr = queue.poll();
                if (curr.left != null) currQueue.offer(curr.left);
                if (curr.right != null) currQueue.offer(curr.right);
                currRes.add(curr.val);
            }

            if (flag) Collections.reverse(currRes);
            flag = !flag;
            res.add(currRes);
            queue = currQueue;
        }

        return res;
    }
}
