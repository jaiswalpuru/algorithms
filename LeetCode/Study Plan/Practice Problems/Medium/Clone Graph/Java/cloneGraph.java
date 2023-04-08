/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;
    public Node() {
        val = 0;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val) {
        val = _val;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Node newNode = new Node(node.val);
        HashMap<Integer, Node> visited = new HashMap<>();
        Queue<Node> q = new ArrayDeque<>();
        visited.put(newNode.val, newNode);
        q.offer(node);

        while(!q.isEmpty()) {
            Node curr = q.poll();
            for (Node nei : curr.neighbors) {
                if (!visited.containsKey(nei.val)) {
                    q.offer(nei);
                    visited.put(nei.val, new Node(nei.val));
                }
                visited.get(curr.val).neighbors.add(visited.get(nei.val));
            }
        }
        return newNode;
    }
}
