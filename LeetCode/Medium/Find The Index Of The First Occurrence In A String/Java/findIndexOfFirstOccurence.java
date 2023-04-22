class Solution {
    public int strStr(String haystack, String needle) {
        int hSize = haystack.length();
        int nSize = needle.length();
        int i=0;

        while(i<hSize) {
            int k = 0;
            int start = i;
        
            while(start<hSize && k<nSize && haystack.charAt(start) == needle.charAt(k)) {
                start++;
                k++;
            }

            if (k == nSize) return i;
            i++;
        }

        return -1;
    }
}
