class Solution {
    public int maxSatisfied(int[] customers, int[] grumpy, int minutes) {
        int sum=0;
        int t =0;
        int size = customers.length;
        for (int i=0; i<size; i++) {
            if (grumpy[i] == 1) {
                t += customers[i];
            }
            if (i >= minutes) {
                t -= customers[i-minutes] * grumpy[i-minutes];
            }
            sum = Math.max(sum, t);
        }
        
        int res=0;
        res += sum;
        for (int i=0; i<size; i++) 
            if (grumpy[i] == 0) res += customers[i];

        return res;
    }
}
