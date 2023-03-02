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
    public TreeNode correctBinaryTree(TreeNode root) {
        ArrayDeque<Pair<TreeNode, TreeNode>> queue = new ArrayDeque<>();
        queue.offer(new Pair(null, root));
        HashMap<Integer, Boolean> visisted = new HashMap<>();

        while (!queue.isEmpty()) {
            int size = queue.size();

            for (int i=0; i<size; i++) {
                Pair<TreeNode,TreeNode> front = queue.poll();
                TreeNode curr = front.getValue();
                if (curr.left != null) {
                    visisted.put(curr.left.val, true);
                    queue.offer(new Pair(curr, curr.left));
                }
                if (curr.right != null) {
                    if (visisted.getOrDefault(curr.right.val, false)) {
                        TreeNode parent = front.getKey();
                        if (parent.left == curr) {
                            parent.left = null;
                        }else {
                            parent.right = null;
                        }
                        break;
                    }
                    visisted.put(curr.right.val, true);
                    queue.offer(new Pair(curr, curr.right));
                }
            }
        }

        return root;
    }
}
