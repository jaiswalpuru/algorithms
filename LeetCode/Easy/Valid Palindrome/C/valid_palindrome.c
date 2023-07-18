bool isPalindrome(char * s){
    int n = strlen(s);
    char* s1 = (char*) malloc(n*sizeof(char));
    int k = 0;
    for (int i=0; i<n; i++)
        if (isdigit(s[i]) || isalpha(s[i]))
            s1[k++] = tolower(s[i]);
    
    for (int i=0; i<k/2; i++)
        if (s1[i] != s1[k-1-i]) return false;
    
    return true;
}
