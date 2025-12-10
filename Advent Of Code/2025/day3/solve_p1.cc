#include <iostream>
#include <fstream>

int solve() {
    std::ifstream in_file;
    in_file.open("input.txt", std::ios::in);
    if (!in_file.is_open()) {
        std::cerr << "error opening file\n";
        exit(1);
    }
    int res = 0;
    std::string line;
    while (getline(in_file, line)) {
        int size = line.size();
        int max_val = 0;;
        for (int i = 0; i < size; i++) {
            for (int j = i + 1; j < size; j++) {
                max_val = std::max(max_val, (((line[i] - '0') * 10) + (line[j] - '0')));
            }
        }
        res += max_val;
    }
    return res;
}

int main(int argc, char *argv[]) {
    std::cout<< solve() << "\n";
    return 0;
}
