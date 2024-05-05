#include <stdio.h>

int main(int argc, char **argv) 
{
    auto int x;
    int m;
    scanf("%d", &m);
    {
        int i; // both m and i can be used in this scope;
        for (i = m; i < m; i++) 
        {
            int i; // redefinition error
        }
    }
    return 0;
}

char *myfunction()
{
    char c[] = "apple";
    return c;
}

int func_name() 
{
    int y;
    return y;
}
