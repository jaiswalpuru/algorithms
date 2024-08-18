class Solution {
public:
    int countSeniors(vector<string>& details) {
        int age_60 = 0;
        string age = "";

        for (int i = 0; i < details.size(); i++) {            
            age = "";
            age += details[i][11];
            age += details[i][12];
            if (std::stoi(age) > 60) age_60++;
        }

        return age_60;
    }
};
