class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        int maxPile = 0;

        for (int pile : piles) {
            maxPile = Math.max(maxPile, pile);
        }

        int l = 1;
        int r = maxPile;

        while(l < r) {
            int mid = l + (r-l)/2;
            if (isPossible(mid, piles, h)) {
                r = mid;
            }else {
                l = mid+1;
            }
        }

        return r;
    }

    public boolean isPossible(int banana, int[] piles, int h) {
        double hrs = 0;
        for (int pile : piles) {
            hrs += Math.ceil((double)pile/banana);
        }

        return hrs<=h;
    }
}
