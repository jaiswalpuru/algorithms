/*
// Definition for a QuadTree node.
class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;

    
    public Node() {
        this.val = false;
        this.isLeaf = false;
        this.topLeft = null;
        this.topRight = null;
        this.bottomLeft = null;
        this.bottomRight = null;
    }
    
    public Node(boolean val, boolean isLeaf) {
        this.val = val;
        this.isLeaf = isLeaf;
        this.topLeft = null;
        this.topRight = null;
        this.bottomLeft = null;
        this.bottomRight = null;
    }
    
    public Node(boolean val, boolean isLeaf, Node topLeft, Node topRight, Node bottomLeft, Node bottomRight) {
        this.val = val;
        this.isLeaf = isLeaf;
        this.topLeft = topLeft;
        this.topRight = topRight;
        this.bottomLeft = bottomLeft;
        this.bottomRight = bottomRight;
    }
};
*/

class Solution {
    public Node construct(int[][] grid) {
        int size = grid.length;
        return recur(0, 0, grid, size);
    }

    public Node recur(int x, int y, int[][] grid, int size) {
        if (isSameValue(x, y, grid, size)) {
            return new Node(grid[x][y]==1, true);
        }else {
            Node root = new Node(false, false);
            root.topLeft = recur(x, y, grid, size/2);
            root.topRight = recur(x, y+size/2, grid, size/2);
            root.bottomLeft = recur(x+size/2, y, grid, size/2);
            root.bottomRight = recur(x+size/2, y+size/2, grid, size/2);
            return root;
        }
    }

    public boolean isSameValue(int x, int y, int[][] grid, int size) {
        for (int i=x; i<x+size; i++) {
            for(int j=y; j<y+size; j++) {
                if (grid[x][y] != grid[i][j]) return false;
            }
        }

        return true;
    }
}
