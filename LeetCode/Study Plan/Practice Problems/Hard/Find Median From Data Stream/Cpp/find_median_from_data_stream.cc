#include <iostream>
#include <queue>

typedef std::vector<int> vi;
typedef std::priority_queue<int> mih;
typedef std::priority_queue<int, vi, greater<int>> mah;

class MedianFinder {
    mih lo;
    mah hi;
public:
    MedianFinder() {}
    
    void addNum(int num) {
        lo.push(num);
        hi.push(lo.top());
        lo.pop();
        if (lo.size() < hi.size()) {
            lo.push(hi.top());
            hi.pop();
        }
    }
    
    double findMedian() {
        return lo.size() > hi.size() ? lo.top() : ((double)lo.top()+hi.top())*0.5;
    }
};

int main() {
    
}
