/**
 *
 * brute force solution
 *
 */

class Solution {
    //refer neetcode for explanation
    Map<Integer, Double> cache;
    public double new21Game(int n, int k, int maxPts) {
        cache = new HashMap<>();
        return dfs(0, n, k, maxPts);
    }

    private double dfs(int score, int n, int k, int maxPts) {
        //base case
        if (score >= k) {
            return score <= n ? 1 : 0;
        }

        //if already in cache
        if (cache.containsKey(score)) return cache.get(score);

        double prob = 0;
        for (int i=1; i<=maxPts; i++) {
            prob += dfs(score+i, n, k, maxPts);
        }
        cache.put(score, prob/maxPts);
        return cache.get(score);
    }
}

/**
 *
 * optimized approach
 *
 */

class Solution {
    public double new21Game(int n, int k, int maxPts) {
        if (k == 0) return 1;
        double windowSum = 0;
        for (int i=k; i<k+maxPts; i++) {
            if (i <= n) windowSum += 1;
        }
        Map<Integer, Double> dp = new HashMap<>();
        for (int i=k-1; i>=0; i--) {
            dp.put(i, windowSum/maxPts);
            double remove = 0;
            if (i + maxPts <= n) {
                remove = dp.getOrDefault(i+maxPts, 1.0);
            }
            windowSum += dp.get(i)-remove;
        }

        return dp.get(0);
    }
}
