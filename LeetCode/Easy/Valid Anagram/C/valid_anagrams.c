// using count of each char
bool isAnagram(char * s, char * t){
    int* chars = (int*) malloc(sizeof(int)*26);
    memset(chars, 0, 26*sizeof(int));
    for (int i=0; i<strlen(s); i++) chars[s[i]-'a']++;
    for (int i=0; i<strlen(t); i++) chars[t[i]-'a']--;
    for (int i=0; i<26; i++) if (chars[i] != 0) return false;
    return true;
}

// using sorting
int cmp(const void* a, const void* b) { return (*(char*)a-*(char*)b); }

bool isAnagram(char * s, char * t){
    if (strlen(s) != strlen(t)) return false;
    qsort(s, strlen(s), sizeof(char), cmp);
    qsort(t, strlen(t), sizeof(char), cmp);
    for (int i=0; i<strlen(s); i++) if (s[i] != t[i]) return false;
    return true;
}
