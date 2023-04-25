class SmallestInfiniteSet {
    TreeSet<Integer> popped;
    TreeSet<Integer> addBack;
    int num = 1;
    public SmallestInfiniteSet() {
        popped = new TreeSet<>();
        addBack = new TreeSet<>();
    }
    
    public int popSmallest() {
        int pop = 0;
        if (addBack.size() > 0) {
            pop = addBack.first();
            addBack.remove(pop);
        } else {
            pop = num;
            num++;
        }
        
        popped.add(pop);
        return pop;
    }
    
    public void addBack(int num) {
        if (popped.contains(num)) popped.remove(num);
        if (this.num > num) addBack.add(num);
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet obj = new SmallestInfiniteSet();
 * int param_1 = obj.popSmallest();
 * obj.addBack(num);
 */
