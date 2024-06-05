#include <stdio.h>
#include <stdlib.h>

#define IS_UPPER(c) \
    (c >= 'A' && c <= 'Z')

#define IS_LOWER(c) \
    (c >= 'a' && c <= 'z')

int main(int argc, char **argv)
{
    printf("is lower %d\n", IS_LOWER('a'));
    printf("is upper %d\n", IS_LOWER('Z'));
    printf("is lower %d\n", IS_LOWER('a'));
    printf("is upper %d\n", IS_UPPER('Z'));
    return 0;
}
