#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int get_decimal_value(long long bin);

int main(int argc, char **argv)
{
    long long bin = 1001001;
    printf("%d \n", get_decimal_value(bin));
    return 0;  
}

int get_decimal_value(long long bin)
{
    int res = 0;
    int i = 0;
    while (bin)
    {
        res += (bin%10) * pow(2, i);
        bin /= 10;
        i++;
    }
    return res;
}
