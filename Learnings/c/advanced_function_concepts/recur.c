#include <stdio.h>

int sum(int n);

int main(int argc, char **argv)
{
    printf("%d", sum(5));
    return 0;
}

int sum(int n) {
    if (n == 0) return n;
    return n + sum(n-1);
}

