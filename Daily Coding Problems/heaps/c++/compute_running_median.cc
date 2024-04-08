#include <iostream>
#include <vector>
#include <queue>

typedef std::vector<int> vi;
typedef std::priority_queue<int> max_pqi;
typedef std::priority_queue<int, vi, std::greater<int>> min_pqi;

double compute_running_median(max_pqi &max_heap, min_pqi &min_heap, int num) {
    min_heap.push(num);
    max_heap.push(min_heap.top());
    min_heap.pop();
    if (min_heap.size() < max_heap.size()) {
        min_heap.push(max_heap.top());
        max_heap.pop();
    }

    if (min_heap.size() > max_heap.size()) {
        return min_heap.top();
    } else {
        return (min_heap.top() + max_heap.top()) * 0.5;
    }
}

int main() {
    vi v = {2, 1, 5, 7, 2, 0, 5};
    max_pqi mxq;
    min_pqi miq;

    for (int i = 0; i < v.size(); i++) {
        std::cout<<compute_running_median(mxq, miq, v[i])<<"\n";
    }
    return 0;    
}
