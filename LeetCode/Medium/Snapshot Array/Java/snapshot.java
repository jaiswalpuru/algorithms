class SnapshotArray {
    int snapId = 0;
    TreeMap<Integer, Integer>[] snaps;
    public SnapshotArray(int length) {
        snaps = new TreeMap[length];
        for (int i=0; i<length; i++) {
            snaps[i] = new TreeMap<>();
            snaps[i].put(0, 0);
        }
    }
    /*
    *   hashmap[]
     */
    public void set(int index, int val) {
        snaps[index].put(snapId, val);
    }
    
    public int snap() {
        return snapId++;
    }
    
    public int get(int index, int snap_id) {
        return snaps[index].floorEntry(snap_id).getValue();
    }
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * SnapshotArray obj = new SnapshotArray(length);
 * obj.set(index,val);
 * int param_2 = obj.snap();
 * int param_3 = obj.get(index,snap_id);
 */
