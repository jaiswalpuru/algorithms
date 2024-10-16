class Solution {
public:
    string longestDiverseString(int a, int b, int c) {
        string res = "";
        priority_queue<pair<int, char>> pq;
        
        if (a > 0) pq.push({a, 'a'});
        if (b > 0) pq.push({b, 'b'});
        if (c > 0) pq.push({c, 'c'});

        while (!pq.empty()) {
            auto [cnt1, char1] = pq.top();
            pq.pop();

            if (res.length() >= 2 && res[res.length() - 1] == char1 && res[res.length() - 2] == char1) {
                if (pq.empty()) break;
                auto [cnt2, char2] = pq.top();
                pq.pop();
                res += char2;
                cnt2--;
                if (cnt2 > 0) pq.push({cnt2, char2});
            } else {
                cnt1--;
                res += char1;
            }

            if (cnt1 > 0) pq.push({cnt1, char1});
        }

        return res;
    }
};
