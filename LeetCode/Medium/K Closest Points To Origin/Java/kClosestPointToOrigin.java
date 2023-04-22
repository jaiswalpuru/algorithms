class Entry {
    Double dis;
    int[] point;
    public Entry(Double d, int[] p) {
        this.dis = d;
        this.point = p;
    }
}

class Solution {
    public int[][] kClosest(int[][] points, int k) {
        int[][] res = new int[k][2];
        PriorityQueue<Entry> pq = new PriorityQueue<>(new Comparator<>(){
            public int compare(Entry a, Entry b) {
                if (a.dis < b.dis) return -1;
                else if (a.dis > b.dis) return 1;
                else return 0;
            }
        });

        HashMap<Double, List<int[]>> disPoint = new HashMap<>();
        for (int[] p : points) disPoint.computeIfAbsent(dis(p), x->new ArrayList<>()).add(p);

        for (Map.Entry<Double, List<int[]>> m : disPoint.entrySet()) {
            List<int[]> coord = m.getValue();
            for (int i=0; i<coord.size(); i++) pq.add(new Entry(m.getKey(), coord.get(i)));
        }
    
        while(!pq.isEmpty() && k-- > 0) res[k] = pq.poll().point;

        return res;
    }

    private double dis(int[] p) {
        return Math.sqrt(p[0]*p[0] + p[1]*p[1]);
    }
}
