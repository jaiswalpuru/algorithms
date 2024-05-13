#include <stdio.h>
#include <stdlib.h>

struct packed_struct {
    unsigned int : 3; // here the variable name is not specified, unnamed fields are used for padding
    unsigned int f1:1;
    unsigned int f2:1;
    unsigned int f3:1;
    unsigned int type:8;
    unsigned int index:18;
};
// c compiler will pack these values together

int main(int argc, char **argv)
{
    struct packed_struct packed_data;
    packed_data.type = 7;
    unsigned int packed_type = packed_data.type;

    unsigned int packed_index = packed_data.index/5 + 1;

    if (packed_data.f2) 
    {
        // do something
    }

    return 0;
}
    
