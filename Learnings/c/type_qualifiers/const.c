#include <stdio.h>
#include <stdlib.h>


const double PI = 3.141592;
const char * MONTHS[12] = {"", "", "", "", "", "", "", "", "", "", "", "", ""};


int main(int argc, char **argv)
{
    const double pi = 3.141592654; // we can put const as many times we want. can also be written as const const const double pi = ..; will mean the same

    typedef const int zip;
    const zip p = 24;


    const float *pf;// pf points to a constant value, the value cannot be changed to which the pointer is pointing to
                    //

    float * const pt; // pt is a constant pointer, address can't change must always point to the same address.
                    // the pointed to value can change but not the address.

    // if * is placed after the const keyword that means we can't change the value the pointer is pointing to
    // if * is placed before the const keyword then the address cannot be changed.
    //
    //

    const float * const ptr; // this means that the pointer must point at the same location and the value stored at the pointer must also not change.
    
    // the order doesn't matter float const or const float
    //

    float const *ptc; // means the same as const float *ptc
                      //



    
    return 0;
}

void display(const int array[], int limit); // here we enforce that array should not be changed by display function.

