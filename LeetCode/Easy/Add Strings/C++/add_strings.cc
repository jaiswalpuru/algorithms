class Solution {
public:
    string addStrings(string num1, string num2) {
        if (num1.length() < num2.length()) return addStrings(num2, num1);

        int i = num2.length()-1;
        int k = num1.length()-1;
        string res = "";
        int c = 0;
        while (i >= 0) {
            int n1=num1.at(k--)-'0';
            int n2=num2.at(i--)-'0';
            int sum = n1+n2+c;
            res = to_string(sum%10)+res;
            c = sum/10;
        }
        while(k >= 0) {
            int sum = num1.at(k--)-'0'+c;
            res = to_string(sum%10) + res;
            c = sum/10;
        }
        if (c > 0)
            res = to_string(c) + res;

        return res;
    }
};
