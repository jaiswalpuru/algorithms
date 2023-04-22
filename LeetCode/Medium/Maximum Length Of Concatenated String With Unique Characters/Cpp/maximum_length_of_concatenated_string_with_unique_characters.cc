#include <iostream>
#include <unordered_map>
#include <vector>

typedef std::vector<std::string> vs;
typedef std::unordered_map<char, int> uci;

bool check_dup(vs v);
void backtrack(vs v, size_t *res, size_t ind, vs temp);
std::string join_string(vs v);

int maximum_length_of_concatenated_string_with_unique_characters(vs &v) {
    size_t res =0;
    backtrack(v, &res, 0, vs{});
    return res;
}

bool check_dup(vs v) {
    uci um;
    for (auto itr = v.begin(); itr!=v.end(); itr++) {
        std::string word = *itr;
        for (int i=0; i<word.size(); i++) {
            um[word[i]]++;
            if (um[word[i]] >1 ) return true;
        } 
    }
    return false;
}

void backtrack(vs v, size_t *res, size_t ind, vs temp) {
    size_t cnt = 0;
    for (int i=ind; i<v.size(); i++){
        temp.push_back(v[i]);
        std::string word = join_string(temp);
        if (!check_dup(temp)) {
            cnt = word.size();
            backtrack(v, res, i+1, temp);
        }
        *res = std::max(*res, cnt);
        temp.pop_back();
    }
}

std::string join_string(vs v) {
    std::string s = "";
    for (auto itr = v.begin(); itr!=v.end(); itr++) s += *itr;
    return s;
}

int main() {
    vs v{"un","iq","ue"};
    std::cout<<maximum_length_of_concatenated_string_with_unique_characters(v)<<"\n";
}
