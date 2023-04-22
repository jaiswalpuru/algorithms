class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        HashMap<Integer, Integer> count = new HashMap<>();
        List<Integer> res = new ArrayList<>();
        for (int val : nums) count.put(val, count.getOrDefault(val, 0)+1);
        
        for (int i=1; i<=nums.length; i++)
            if (!count.containsKey(i)) res.add(i);
        
        return res;
    }
}
