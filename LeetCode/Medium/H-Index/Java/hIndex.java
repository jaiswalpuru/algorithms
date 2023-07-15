class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);
        int i = 0;
        int n = citations.length;
        while(i < n && citations[n-1-i] > i) i++;
        return i;
    }
}
