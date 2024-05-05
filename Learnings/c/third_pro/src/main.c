#include <stdio.h>
#include <stdlib.h>

double* multiply_by_two(double* in) 
{
    double *twice = malloc(sizeof(double));
    *twice = *in * 2.0;
    return twice; 
}

int main(int argc, char **argv) 
{
    int *age = malloc(sizeof(int));
    *age = 30;
    double *salary = malloc(sizeof(double));
    *salary = 123456.8;
    double *my_list = malloc(3 * sizeof(double));
    my_list[0] = 1.0;
    my_list[1] = 2.0;
    my_list[2] = 3.0;
    double *my_twice_sal = multiply_by_two(salary);
    printf("double salary : %0.3f", *my_twice_sal);
    free(age);
    free(salary);
    free(my_list);
    free(my_twice_sal);
    return 0;
}
