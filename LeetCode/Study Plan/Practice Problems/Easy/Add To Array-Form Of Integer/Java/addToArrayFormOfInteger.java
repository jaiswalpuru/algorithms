class Solution {
    public List<Integer> addToArrayForm(int[] num, int k) {
        List<Integer> res = new ArrayList<>();
        int size = num.length;
        int i = size-1;
        int carry = 0;

        while (i >= 0 && k > 0) {
            int sum = num[i] + k%10 + carry;
            res.add(sum%10);
            carry = sum/10;
            i--;
            k/=10;
        }

        while(i >= 0) {
            int sum = num[i] + carry;
            res.add(sum%10);
            carry = sum/10;
            i--;
        }

        while(k > 0) {
            int sum = k%10 + carry;
            res.add(sum%10);
            carry = sum/10;
            k/=10;
        }
        
        while(carry > 0) {
            res.add(carry%10);
            carry/=10;
        }

        Collections.reverse(res);
        return res;
    }
}
