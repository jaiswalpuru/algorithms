#include <stdio.h>
#include <stdlib.h>

union car {
    int i_val;
    float f_val;
};// here it can also be ended as } car1, car2, *car3 ;

union car_2 {
    int i_val;
    float f_val;
    char c_val[40];
};

struct owner {
    char ssn[12];
};

struct lease_company {
    char name[40];
    char headquarters[40];
};

struct car_data {
    char make[15];
    int status;

    //below is a anonymous union
    union {
        struct owner owncar;
        struct lease_company lease_car;
    };
};


// this is setting up an array of table, each array member of struct->data can contain either int  float or char
struct {
    char *name;
    enum symbol_type type;
    union{
        // entries inside union
        int i;
        float f;
        char c;
    }; data;
} table[entries];

int main(int argc, char **argv) 
{
    union car car1, car2, *car3;
    union car_2 ca;
    printf("size occupied by car : %zu\n", sizeof(car1));
    printf("size occupied by car : %zu\n", sizeof(car3));
    printf("size occupied by car2 : %zu\n", sizeof(ca));
    return 0;    
}
