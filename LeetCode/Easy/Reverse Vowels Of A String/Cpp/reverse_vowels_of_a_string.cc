#include <iostream>
#include <vector>

typedef std::vector<char> vc;

std::string reverse_vowels_of_a_string(std::string word) {
    vc vow;
    int k =0;
    for (int i=0; i<word.length(); i++)
        if (word[i] =='a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] =='u' || word[i]=='A' || word[i]== 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U') vow.push_back(word[i]);
    reverse(vow.begin(), vow.end());
    for (int i=0;i<word.length();i++) {
        if (word[i] =='a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] =='u' || word[i]=='A' || word[i]== 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U') {
            word[i] = vow[k];
            k++;
        }
    }
    return word;
}

int main () {
    std::cout<<reverse_vowels_of_a_string("hello")<<"\n";
}
