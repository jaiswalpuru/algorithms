#include <iostream>
#include <vector>
#include <cstring>
#include <unordered_map>

typedef std::vector<std::string> vs;
typedef std::unordered_map<std::string, bool> usb;

int unique_email_addresses(vs email) {
    usb visited;
    for (int i=0; i<email.size(); i++) {
        std::string curr = email[i];
        std::string temp = "";
        bool plus = false;
        bool rate = false;
        int j = 0;
        while(j<curr.length()) {
            if (!rate) {
                if (curr[j] != '@') {
                    if (plus) {
                        j++;
                        continue;
                    }else {
                        if (curr[j] == '+') {
                            j++;
                            plus = true;
                            continue;
                        }else {
                            if (curr[j] != '.') temp+=curr[j];
                        }
                    }
                }else {
                    temp += curr[j];
                    rate = true;
                }
            }else {
                temp += curr[j];
            }
            j++;
        }
        visited[temp] = true;
    }
    return visited.size();
}

int main() {
    std::cout<<unique_email_addresses(vs{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"})<<"\n";
}
