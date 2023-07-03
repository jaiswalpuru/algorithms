bool buddyStrings(char * s, char * goal){
    int n = strlen(s);
    int m = strlen(goal);
    if (n != m) return false;
    if (strcmp(s, goal)==0) {
        int *freq = (int *)malloc(26*sizeof(int));
        memset(freq, 0, 26);
        for (int i=0; i<n; i++) {
            freq[s[i]-'a']++;
            if (freq[s[i]-'a'] == 2) return true;
        }
        return false;
    }

    int first_index = -1;
    int second_index = -1;
    for (int i=0; i<n; i++) {
        if (s[i] != goal[i]) {
            if (first_index == -1) first_index = i;
            else if (second_index == -1) second_index = i;
            else return false;
        }
    }

    if (second_index == -1) return false;
    return s[first_index] == goal[second_index] && s[second_index] == goal[first_index];
}
