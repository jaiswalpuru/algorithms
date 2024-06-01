#define pdv pair<double, vector<int>>
typedef double d;
typedef vector<int> vi;
typedef vector<pdv> vpdv;

class Solution {
public:
    vector<int> kthSmallestPrimeFraction(vector<int>& arr, int k) {
        priority_queue<pdv, vpdv, greater<pdv>> pq;
        
        for (int i = 0; i < arr.size(); i++) {
            for (int j = 0; j < arr.size(); j++) {
                if (i == j) continue;
                pq.push(make_pair((d)arr[i]/(d)arr[j], vector<int>{arr[i], arr[j]}));
            }
        }

        while (!pq.empty() && k > 1) {
            pq.pop();
            k--;
        }

        return pq.top().second;
    }
};
