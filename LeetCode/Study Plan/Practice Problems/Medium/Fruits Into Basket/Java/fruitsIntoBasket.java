class Solution {
    public int totalFruit(int[] fruits) {
        HashMap<Integer, Integer> baskets = new HashMap();
        int k = 0;
        int size = fruits.length;
        int maxFruits = 0;

        for (int i=0; i<size; i++) {
            baskets.put(fruits[i], baskets.getOrDefault(fruits[i], 0)+1);
            while (baskets.size() > 2) {
                baskets.put(fruits[k], baskets.get(fruits[k])-1);
                if (baskets.get(fruits[k])==0) baskets.remove(fruits[k]);
                k++;
            }
            maxFruits = Math.max(maxFruits, i-k+1);
        }
        
        return maxFruits;
    }
}
