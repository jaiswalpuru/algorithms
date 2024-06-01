#include <stdio.h>

//#define JUST_CHECKING
#define LIMIT 4

#define MYDEF 4


#define US 0
#define UK 1
#define FRANCE 2
#define GERMANY 3
#define COUNTRY US

int main(int argc, char **argv)
{
    int i;
    int total = 0;

    for (i = 1; i < LIMIT; i++) {
        total += 2 * i * i + 1;
        
#ifdef JUST_CHECKING
        printf(" i = %d, running total = %d \n ", i, total);
#endif
    }

    printf ("grand total = %d\n", total);

#if MYDEF == 5
    printf("mydef has value of 5\n");
#endif

#if MYDEF == 4
    printf("mydef has value of 4\n");
#endif

#if MYDEF == 4 && LIMIT == 4
    printf("mydef has value of 4 and limit is also 4\n");
#endif

#if COUNTRY == US || COUNTRY == UK
#define GREETING "Hello."
#elif COUNTRY == FRANCE
#define GREETING "Bonjour"
#endif

    printf("%s", GREETING);

    return 0;
}
