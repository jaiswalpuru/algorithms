#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

double avg (double v1, double v2, ...);

int main (int argc, char **argv)
{
    double v1 = 10.5, v2 = 2.5;
    int num1 = 6, num2 = 5;
    long num3 = 12L, num4 = 20L;

    printf("average  : %.21f\n", avg(v1, v2, 3.5, 4.5, 0.0));
    return 0;
}

// assumtion we are making here is that the last param is always 0.0 here for the loop

double avg (double v1, double v2, ...)
{
    va_list parg;
    double sum = v1 + v2;
    int cnt = 2;
    va_start(parg, v2);
    double val = 0.0;

    while((val = va_arg(parg, double)) != 0.0) {
        sum += val;
        ++cnt;
    }
    va_end(parg);

    return sum / cnt;
}
    
