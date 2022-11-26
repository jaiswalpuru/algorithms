#include <iostream>
#include <vector>

typedef std::vector<std::string> vs;

int shortest_word_distance(vs words_dict, std::string word1, std::string word2) {
    int ind1 = -1, ind2 = -1;
    int dis = INT_MAX;
    for (int i=0; i<words_dict.size(); i++) {
        if (words_dict[i].compare(word1) == 0) ind1 = i;
        if (words_dict[i].compare(word2) == 0) ind2 = i;
        if (ind1 != -1 && ind2 != -1) dis = std::min(dis, abs(ind1-ind2));
    }
    return dis;
}

int main() {
    std::cout<<shortest_word_distance(vs{"practice", "makes", "perfect", "coding", "makes"}, "practice", "coding")<<"\n";
}
