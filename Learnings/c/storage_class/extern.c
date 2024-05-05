#include <stdio.h>

int i;
char text[255];
char t[255][50];
int count;

void foo(void);
extern void write_extern();

int main(int argc, char **argv)
{
    printf("%i \n", i);
    foo();
    printf("%i \n", i);
    printf("%s \n", text);
    count = 5;
    write_extern();
    return 0;
}
