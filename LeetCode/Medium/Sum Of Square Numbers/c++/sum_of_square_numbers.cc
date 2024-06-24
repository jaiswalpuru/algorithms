typedef long long ll;

class Solution {
public:
    bool judgeSquareSum(int c) {
        ll hi = sqrt(c);
        ll lo = 0;
        ll val = 0;

        if (c == 1 || c == 0) return true;

        while (lo <= hi) {
            val = lo * lo + hi * hi;
            if (val == c) return true;
            else if (val > c) hi--;
            else lo++;
        }
        
        return false;
    }
};
