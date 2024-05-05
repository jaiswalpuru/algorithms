#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern int i;
extern char text[];
extern char t[][50];
extern int count;

void write_extern(void)
{
    printf("count %d\n", count);
}

void foo(void)
{
    i = 200;
    strcpy(text, "Hell");
}
