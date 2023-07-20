class Solution {
    public String convertToBase7(int num) {
        if (num == 0) return "0";
        boolean isNeg = num < 0;
        if (isNeg) num *= -1;
        StringBuilder sb = new StringBuilder();
        while(num > 0) {
            sb.append(num%7);
            num = num/7;
        }
        return isNeg ? "-" + sb.reverse().toString() : sb.reverse().toString();
    }
}
