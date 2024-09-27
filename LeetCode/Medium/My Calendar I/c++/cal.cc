class MyCalendar {
public:
    set<pair<int, int>> cal;
    MyCalendar() {}
    
    bool book(int start, int end) {
        const pair<int, int> event{start, end};
        auto next_eve = cal.lower_bound(event);
        if (next_eve != cal.end() && next_eve->first < end) return false;
        if (next_eve != cal.begin()) {
            const auto prev_eve = prev(next_eve);
            if (prev_eve->second > start) return false;
        }
        cal.insert(event);
        return true;
    }
};

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar* obj = new MyCalendar();
 * bool param_1 = obj->book(start,end);
 */
