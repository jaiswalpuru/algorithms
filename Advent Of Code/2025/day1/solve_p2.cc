#include <iostream>
#include <fstream>

int solve() {
    std::ifstream in_file;
    in_file.open("input2.txt", std::ios::in);
    if (!in_file.is_open()) {
        std::cerr << "error opening file\n";
        exit(1);
    }
    std::string line;
    int start = 50;
    int password = 0;
    while (std::getline(in_file, line)) {
        int dir = line[0] == 'L' ? -1 : 1;
        int val = stoi(line.substr(1));
        for (int i = 0; i < val; i++) {
            start += dir;
            start %= 100;
            if (start == 0) password++;
        }
    }
    in_file.close();

    return password;
}

int main(int argc, char *argv[]) {
    std::cout << solve() << "\n";
    return 0;
}
