#include <stdio.h>

int sum2d(int, int, int arr[*][*]); // the names here can be ommiited and can we written sum2d(int r, int c, int arr[*][*]) only with c99 and c11

#ifdef __STDC_NO_VLA__
    printf("Not supported\n");
    exit(1);
#endif

int main(int argc, char **argv)
{
    size_t size = 0;
    printf("enter the number of elemets you want to store : ");
    scanf("%zd", &size);
    float arr[size];
    
    size_t rows = 0, cols = 0;
    
    printf("Enter rows and cols : ");
    scanf("%zd %zd", &rows, &cols);
    float aarr[rows][cols];
    

    return 0;
} 


int sum2d(int r, int c, int arr[r][c])
{
    int i = 0;
    int j = 0;
    int sum = 0;

    for (; i < r; i++)
    {
        for (; j < c; j++)
        {
            sum += arr[i][j];
        }
    }
    return sum;
}
