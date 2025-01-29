#include <iostream>
#include <queue>
#include <unordered_map>

using namespace std;

vector<string> get_k_most_frequeunt_strings(vector<string> strs, int k) {
    unordered_map<string, int> cnt;

    priority_queue<pair<int, string>> pq;
    for (string s : strs) cnt[s]++;

    for (auto p : cnt) pq.push(make_pair(p.second, p.first));

    vector<string> res;
    while (k-- > 0) {
        auto curr = pq.top();
        pq.pop();
        res.push_back(curr.second);
    }
    return res;
}

int main(int argc, char** argv) {
    vector<string> res = get_k_most_frequeunt_strings({"go", "coding", "byte", "byte", "go", "interview", "go"}, 2);
    for (string s : res) cout << s << " ";
    cout << "\n";
    return 0;
}
