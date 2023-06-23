class Solution {
    public List<List<Integer>> findWinners(int[][] matches) {
        HashMap<Integer, Integer> win = new HashMap<>();
        HashMap<Integer, Integer> lose = new HashMap<>();

        for (int[] m : matches) {
            win.put(m[0], win.getOrDefault(m[0], 0) + 1);
            lose.put(m[1], lose.getOrDefault(m[1], 0) + 1);
        }

        List<List<Integer>> res = new ArrayList<>();
        List<Integer> r1 = new ArrayList<>();
        List<Integer> r2 = new ArrayList<>();

        for (Integer key : win.keySet())
            if (!lose.containsKey(key)) r1.add(key);
        
        for (Integer key : lose.keySet()) 
            if (lose.get(key) == 1) r2.add(key);

        Collections.sort(r1);
        Collections.sort(r2);
        res.add(r1);
        res.add(r2);
        return res;
    }
}
