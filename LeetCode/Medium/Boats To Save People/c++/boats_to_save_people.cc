class Solution {
public:
    int numRescueBoats(vector<int>& people, int limit) {
        int l = 0, h = people.size()-1;
        int sum = 0;
        int min_boats = 0;

        sort(people.begin(), people.end());

        while (l <= h) {
            sum = people[l] + people[h];
            if (sum <= limit) {
                l++;
            } 
            h--;
            min_boats++;
        }

        return min_boats;
    }
};
