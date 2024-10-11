class Solution {
public:
    int smallestChair(vector<vector<int>>& times, int targetFriend) {
        vector<pair<int, int>> events;
        priority_queue<int, vector<int>, greater<int>> avail;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> occ;

        for (int i = 0; i < times.size(); i++) {
            events.push_back({times[i][0], i});
            events.push_back({times[i][1], -i});
        }

        sort(events.begin(), events.end());
        for (int i = 0; i < times.size(); i++) avail.push(i);

        for (auto& event : events) {
            int time = event.first;
            int ind = event.second;

            while (!occ.empty() && occ.top().first <= time) {
                avail.push(occ.top().second);
                occ.pop();
            }

            if (ind >= 0) {
                int chair = avail.top();
                avail.pop();
                if (ind == targetFriend) return chair;
                
                occ.push({times[ind][1], chair});
            }
        }

        return 0;
    }
};
