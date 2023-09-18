class Solution {
    public int[] kWeakestRows(int[][] mat, int k) {
        PriorityQueue<Map.Entry<Integer, Integer>> pq = new PriorityQueue<>(new Comparator<>() {
            public int compare(Map.Entry<Integer, Integer> a, Map.Entry<Integer, Integer> b) {
                if (a.getValue() == b.getValue()) return a.getKey()-b.getKey();
                return a.getValue()-b.getValue();
            }
        });
        for (int i=0; i<mat.length; i++) {
            int sol = 0;
            for (int j=0; j<mat[0].length; j++) {
                if (mat[i][j] == 1) sol++;
            }
            pq.offer(Map.entry(i, sol));
        }
        int[] res = new int[k];
        for (int i=0; i<k; i++) {
            res[i] = pq.poll().getKey();
        }
        return res;
    }
}
