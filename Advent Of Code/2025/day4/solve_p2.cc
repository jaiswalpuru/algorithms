#include <iostream>
#include <vector>
#include <fstream>

int solve() {
    std::ifstream in_file;
    in_file.open("input.txt", std::ios::in);
    if (!in_file.is_open()) {
        std::cerr << "error opening file\n";
        exit(1);
    }
    std::string line;
    std::vector<std::vector<char>> vec;
    while (getline(in_file, line)) {
        std::vector<char> col_vec;
        for (int i = 0; i < line.size(); i++) col_vec.push_back(line[i]);
        vec.push_back(col_vec);
    }
    int m = vec.size();
    int n = vec[0].size();
    int num_ways = 0;
    std::vector<std::vector<int>> dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}};
    bool can_be_removed = true;
    while (can_be_removed) { 
        can_be_removed = false;
        for (int i = 0; i < vec.size(); i++) {
            for (int j = 0; j < vec[i].size(); j++) {
                if (vec[i][j] == '@') {
                    int rolls = 0;
                    for (int d = 0; d < 8; d++) {
                        int x = dir[d][0] + i;
                        int y = dir[d][1] + j;
                        if (x >= m || y >= n || x < 0 || y < 0) continue;
                        if (vec[x][y] == '@') rolls++;
                    }
                    if (rolls < 4) {
                        num_ways++;
                        can_be_removed = true;
                        vec[i][j] = '.';
                        break;
                    }
                }
            }
        }
    }
    return num_ways;
}

int main(int argc, char *argv[]) {
    std::cout << solve() << "\n";
    return 0;
}
