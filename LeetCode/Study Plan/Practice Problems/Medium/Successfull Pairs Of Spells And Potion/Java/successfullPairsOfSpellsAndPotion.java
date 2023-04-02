class Solution {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        int n = spells.length;
        int m = potions.length;
        Arrays.sort(potions);
        int[] res = new int[n];
        for (int i=0; i<n; i++) {
            res[i] = bSearch(spells[i], potions, success);
        }
        return res;
    }

    private int bSearch(long spell, int[] potions, long success) {
        int low = 0; 
        int high = potions.length-1;
        int ind = -1;

        while(low <= high) {
            int mid = low + (high-low)/2;
            long val = potions[mid] * spell;
            if (val >= success) {
                high = mid-1;
                ind = mid;
            }else {
                low = mid+1;
            }
        }
        if (ind == -1) return 0;
        return potions.length-ind;
    }
}
