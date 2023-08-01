bool isValid(char * s){
    int size = strlen(s);
    char stack[size];
    int top = -1;
    for (int i=0; i<size; i++) {
        char ch = s[i];
        if (ch == ')' || ch == ']' || ch == '}') {
            if (top == -1) return false;
            if ((ch == ')' && stack[top] != '(') || (ch == ']' && stack[top] != '[') || 
            (ch == '}' && stack[top] != '{')) return false;
            top--;
        } else {
            stack[++top] = ch;
        }
    }
    return top == -1 ? true : false;
}
