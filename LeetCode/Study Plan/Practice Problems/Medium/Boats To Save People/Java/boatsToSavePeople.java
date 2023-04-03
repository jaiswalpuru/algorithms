class Solution {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int low=0, high = people.length-1;
        int cnt = 0;
        while(low <= high) {
            int total = people[low] + people[high];
            if (total <= limit) {
                low++;
                high--;
            }else {
                high--;
            }
            cnt++;
        }
        return cnt;
    }
}
