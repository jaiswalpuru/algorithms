class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        int freq[26] = {0};
        priority_queue<int> pq;
        for (char t : tasks) freq[t - 'A']++;
        for (int i = 0; i < 26; i++) {
            if (freq[i] > 0) pq.push(freq[i]);
        }

        int time = 0;
        while (!pq.empty()) {
            int cycle = n + 1;
            vector<int> store;
            int task_cnt = 0;
            while (cycle-- && !pq.empty()) {
                if (pq.top() > 1) store.push_back(pq.top() - 1);
                pq.pop();
                task_cnt++;
            }

            for (int& x : store) pq.push(x);
            time += (pq.empty() ? task_cnt : n + 1);
        }
        return time;
    }
};
