#include <iostream>
#include <fstream>

int solve() {
    std::ifstream in_file;
    in_file.open("input.txt", std::ios::in);
    if (!in_file.is_open()) {
        std::cerr << "error opening file \n";
        exit(1);
    }
    std::string line;
    int start = 50;
    int password = 0;
    while (std::getline(in_file, line)) {
        int dir = line[0] == 'L' ? 0 : 1;
        int val = (std::stoi(line.substr(1))) % 100;
        if (dir) {
            start = (start + val) % 100;
        } else {
            val = 100 - val;
            start = (start + val) % 100;
        }
        if (start == 0) password++;
    }

    in_file.close();
    return password;
}

int main(int argc, char *argv[]) {
    std::cout << solve() << "\n";
    return 0;   
}
