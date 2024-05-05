#include <stdio.h>

int fun() 
{
    static int cnt;
    int local_var = 0;
    printf("auto %d static %d \n", cnt, local_var);
    local_var++;
    cnt++;
    return cnt;
}

int main(int argc, char **argv) 
{
    printf("%d \n", fun());
    printf("%d \n", fun());
    return 0;
}
