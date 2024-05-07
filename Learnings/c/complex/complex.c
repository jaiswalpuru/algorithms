#include <stdio.h>
#include <complex.h>

int main(int argc, char **argv)
{
#ifdef __STDC_NO_COMPLEX__
    printf("complex numbers not supported\n");
    exit(0);
#else
    printf("complex numbers are supported\n");
    double complex cx = 1.0 + 3.0*I;
    double complex cy = 1.0 - 4.0*I;

    printf("working with complex numbers : \n");
    printf("starting values : cx = %.2f + %.2f cy = %.2f + %.2f \n", creal(cx), cimag(cx), creal(cy), cimag(cy));

    double complex sum = cx + cy;
    printf("sum : sum = %.2f + %.2f \n", creal(sum), cimag(sum));
    
    double complex conjugate = conj(sum); // can also define as double _Complex if not including comlpex.h
    printf("conj = %.2f + %.2f \n", creal(conjugate), cimag(conjugate));
    
    //complex exponentiation
    double complex pwr = cpow(cx, cy);
    printf("complex power = %.2f + %.2f \n", creal(pwr), cimag(pwr));

#endif

    return 0;
}
    
