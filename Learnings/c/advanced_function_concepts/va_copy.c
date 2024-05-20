#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <math.h>

double sample_stdv(int cnt, ...);

int main (int argc, char **argv)
{
    printf("%f\n", sample_stdv(4, 25.0, 27.3, 26.9, 25.7));

    return 0;
}


double sample_stdv(int cnt, ...)
{
    double sum = 0;
    va_list args1;
    va_list arg_copy;
    va_start(args1, cnt);
    va_copy(arg_copy, args1);

    for (int i = 0; i < cnt; i++) {
        double num = va_arg(args1, double);
        sum += num;
    }

    va_end(args1);

    double mean = sum/cnt;
    double sum_sq_diff;

    for (int i = 0; i < cnt; i++) {
        double num = va_arg(arg_copy, double);
        sum_sq_diff += (num-mean) * (num-mean);
    }

    va_end(arg_copy);
    return sqrt(sum_sq_diff/cnt);

}
