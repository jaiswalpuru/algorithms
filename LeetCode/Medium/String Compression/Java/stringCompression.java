class Solution {
    public int compress(char[] chars) {
        int i = 0;
        int j = 0;
        int size = chars.length;

        while(i < size) {
            int cnt = 0;
            chars[j] = chars[i];
            while(i<size && chars[i] == chars[j]) {
                cnt++;
                i++;
            }

            if (cnt > 1) {
                String val = ""+cnt;
                for (char c:val.toCharArray()) {
                    j++;
                    chars[j] = c;
                }
            }
            j++;
        }

        return j;
    }
}
