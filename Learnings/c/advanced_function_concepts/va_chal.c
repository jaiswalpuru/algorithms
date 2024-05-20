#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

int sum(int cnt, ...);

int main(int argc, char **argv)
{
    printf("%d\n", sum(4, 10, 20, 30, 40));
    return 0;
}

int sum(int cnt, ...)
{
    int i = 0;
    int sum = 0;
    int res = 0;
    va_list arg;
    va_start(arg, cnt);

    for (i = 0; i < cnt; i++) {
        int num = va_arg(arg, int);
        res += num;
    }

    va_end(arg);
    return res;
}
