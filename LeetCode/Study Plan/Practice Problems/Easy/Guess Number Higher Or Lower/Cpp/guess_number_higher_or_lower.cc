#include <iostream>

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */
typedef long long ll;

class Solution {
public:
    int guessNumber(int n) {
        ll l=0, r=n;
        while(l<=r) {
            ll g = (l+r)/2;
            int guess_val = guess(g);
            if (guess_val == 0) return g;
            if (guess_val == -1) r = g-1;
            else l = g+1;
        }
        return -1;
    }
};
