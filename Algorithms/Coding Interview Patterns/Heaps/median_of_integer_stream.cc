#include <iostream>
#include <queue>

using namespace std;

vector<double> get_data(vector<pair<string, int>> query) {
    priority_queue<int> max_pq;
    priority_queue<int, vector<int>, greater<int>> min_pq;

    vector<double> res;

    for (auto p : query) {
        if (p.first == "add") {
            min_pq.push(p.second);
            max_pq.push(min_pq.top());
            min_pq.pop();
            if (min_pq.size() < max_pq.size()) {
                min_pq.push(max_pq.top());
                max_pq.pop();
            }
        } else {
            double d = 0.0;
            if (min_pq.size() > max_pq.size()) {
                d = (double)min_pq.top();
            } else {
                d = ((double)min_pq.top() + (double)max_pq.top()) * 0.5;
            }
            res.push_back(d);
        }
    }
    return res;
}

int main(int argc, char** argv) {
    vector<double> res = get_data({{"add", 3}, {"add", 6}, {"getmedian", 0}, {"add", 1}, {"getmedian", 0}});
    for (double v : res) cout << v << " ";
    cout << "\n";
    return 0;
}
