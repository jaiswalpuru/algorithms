class Solution {
public:
    int passThePillow(int n, int time) {
        int full_round = time / (n-1);
        time = time % (n-1);
        
        if (full_round % 2 == 0) {
            return time + 1;
        }
        return n - time;
    }
};
