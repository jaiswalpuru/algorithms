#include <stdio.h>

int main(int argc, char **argv)
{
#ifdef MY_MACRO
    printf("macro is defined \n");
#endif

    char c = -10;
    printf("%d", c);

    return 0;
}
