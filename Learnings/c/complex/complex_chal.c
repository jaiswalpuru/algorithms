#include <stdio.h>
#include <math.h>
#include <complex.h>

int main(int argc, char **argv)
{
#ifdef __STDC_NO_COMLPEX__
    printf("complex numbers not supported\n");
    exit(0);
#else
    printf("complex number are supported\n");
    double complex z1 = I * I;
    printf("I * I = %.1f%+.1fi\n", creal(z1), cimag(z1));

    double _Complex z2 = pow(I, 2);
    printf("pow(I, 2) = %.1f%+.1fi\n", creal(z1), cimag(z1));

    double PI = acos(-1);
    double complex z3 = exp(I * PI);

    printf("exp(I * PI) = %.1f%+.1fi\n", creal(z3), cimag(z3));

    double complex z4 = 1+2*I, z5 = conj(z4); // 1-2*I
    double complex z6 = z4 * z5;
    printf("z4 * z5 = %.1f%+.1fi\n", creal(z6), cimag(z6));

#endif
    return 0;
}
