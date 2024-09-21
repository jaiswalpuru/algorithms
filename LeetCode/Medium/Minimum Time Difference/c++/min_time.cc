class Solution {
public:
    int findMinDifference(vector<string>& timePoints) {
        int min_time = INT_MAX;
        vector<int> time_min;

        for (string time : timePoints)
            time_min.push_back(get_minutes(time));

        sort(time_min.begin(), time_min.end());
        time_min.push_back(time_min[0] + get_minutes("24:00"));
        for (int i = 1; i < time_min.size(); i++) {
            min_time = min(min_time, time_min[i]- time_min[i-1]);
        }

        return min_time;
    }

    int get_minutes(string time) {
        return (((time[0] - '0') * 10 + (time[1] - '0')) * 60 + ((time[3]-'0') * 10 + (time[4] - '0')));
    }
};
