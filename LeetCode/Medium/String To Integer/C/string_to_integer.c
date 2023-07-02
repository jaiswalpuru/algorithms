int myAtoi(char * s){
    int is_neg = 1;
    int len = strlen(s);
    int i=0;
    while(s[i] == ' ') i++;
    if (s[i] == '-') {
        is_neg = -1;
        i++;
    }else if (s[i] == '+') i++;
    while(s[i] == '0') i++;
    long long res = 0;
    while(i < len && isdigit(s[i]) && res <= INT_MAX) {
        res = (res*10) + (s[i]-'0');
        i++;
    }
    res *= is_neg;
    if (res > INT_MAX) return INT_MAX;
    if (res < INT_MIN) return INT_MIN;
    return res;
}
