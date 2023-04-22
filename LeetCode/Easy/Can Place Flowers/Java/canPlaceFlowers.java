class Solution {
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        int size = flowerbed.length;

        for (int i=0; i<size; ) {
            if (n == 0) break;
            if (flowerbed[i] == 1) {
                i+=2;
            }else {
                if (i+1 < size && flowerbed[i+1] == 1) {
                    i+=3;
                }else {
                    flowerbed[i] = 1;
                    i+=2;
                    n--;
                }
            }
        }

        return n == 0;
    }
}
