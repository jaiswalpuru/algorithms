#include <iostream>
#include <string>
#include <queue>
#include <set>

using namespace std;

bool explore_lvl(queue<string>& q, set<string>& visited, set<string>& other_visited, set<string> dict) {
    string lower_case = "abcdefghijklmnopqrstuvwxyz";
    int size = q.size();
    for (int k = 0; k < size; k++) {
        string curr = q.front();
        q.pop();
        for (int i = 0; i < curr.length(); i++) {
            for (int j = 0; j < 26; j++) {
                string new_str = curr.substr(0, i) + lower_case[j] + curr.substr(i + 1, curr.length());
                if (other_visited.count(new_str) == 1) return true;
                if (dict.count(new_str) == 1 && visited.count(new_str) == 0) {
                    visited.insert(new_str);
                    q.push(new_str);
                }
            }
        }
    }
    return false;
}
    
int shortest_transformation_bidirectional(string s, string e, vector<string> dict) {
    if (s == e) return 1;
    set<string> dict_set(dict.begin(), dict.end());
    if (dict_set.count(s) == 0 || dict_set.count(e) == 0) return 0;

    set<string> visited_s;
    set<string> visited_e;
    queue<string> q_s;
    queue<string> q_e;
    int lvl_s = 0;
    int lvl_e = 0;
    q_s.push(s);
    q_e.push(e);

    while (!q_s.empty() && !q_e.empty()) {
        lvl_s++;
        if (explore_lvl(q_s, visited_s, visited_e, dict_set)) return lvl_s + lvl_e + 1;

        lvl_e++;
        if (explore_lvl(q_e, visited_e, visited_s, dict_set)) return lvl_s + lvl_e + 1;
    }

    return 0;
}

int shortest_transformation(string s, string e, vector<string> dict) {
    set<string> set_str;
    set<string> visited;
    for (auto str : dict) set_str.insert(str);
    if (set_str.count(s) == 0 || set_str.count(e) == 0) return 0;
    if (s == e) return 1;
    int dist = 0;
    string lower_case = "abcdefghijklmnopqrstuvwxyz";

    queue<string> q;
    q.push(s);
    visited.insert(s);

    while (!q.empty()) {
        int size = q.size();
        for (int i = 0; i < size; i++) {
            string curr = q.front();
            q.pop();
            if (curr == e) return dist + 1;
            for (int k = 0; k < curr.length(); k++) {
                for (int j = 0; j < 26; j++) {
                    string new_str = curr.substr(0, k) + lower_case[j] + curr.substr(k + 1, curr.length());
                    if (visited.count(new_str) == 0 && set_str.count(new_str) == 1) {
                        visited.insert(new_str);
                        q.push(new_str);
                    }
                }
            }
        }
        dist++;
    }
    return 0;
}

int main(int argc, char** argv) {
    string start = "red";
    string end = "hit";
    vector<string> dict = {"red", "bed", "hat", "rod", "rad", "rat", "hit", "bad", "bat"};
    cout << shortest_transformation_bidirectional(start, end, dict) << "\n";
    return 0;
}
