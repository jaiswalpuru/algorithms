class Solution {
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int max = 0;
        List<Boolean> res = new ArrayList<>();

        for (int val : candies) max = Math.max(max, val);
        for (int val : candies) {
            if (val+extraCandies >= max) res.add(true);
            else res.add(false);
        }

        return res;
    }
}
