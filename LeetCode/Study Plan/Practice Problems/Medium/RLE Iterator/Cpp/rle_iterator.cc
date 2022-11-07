#include <iostream>
#include <vector>

typedef std::vector<int> vi;

class RLEIterator {
    vi v;
    RLEIterator(vector<int>& encoding) {
        this->v = encoding;
    }
    
    int next(int n) {
        if (this->v.size() > 0) {
			for (int i=0; i<this->v.size(); i+=2) {
				int cnt = this->v[i];
				int val = this->v[i+1];
				if (cnt >= n) {
					this->v[i] -= n;
					return val;
				}else {
					n-= this->v[i];
					this->v[i] = 0;
				}
			}
        }
		return -1;
    }
};
