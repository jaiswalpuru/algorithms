class Solution {
    public int maxSatisfied(int[] customers, int[] grumpy, int minutes) {
        int l=0;
        int r=0;
        int gCnt=0;
        int sum=0;
        int size = customers.length;
        for (int i=0; i<size; i++) {
            int j = i;
            int temp = 0;
            while(j < Math.min(i+minutes, size)) {
                if (grumpy[j] == 1){
                    temp+=customers[j];
                }
                j++;
            }

            if (temp > sum) {
                sum = temp;
                l=i;
                r=j;
            }
        }
        
        int res=0;
        res += sum;
        for (int i=0; i<size; i++) 
            if (grumpy[i] == 0) res += customers[i];

        return res;
    }
}
