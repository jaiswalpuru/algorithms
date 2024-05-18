#include <stdio.h>
#include <ctype.h>

//int getc(FILE *stream); 
//returns an integer value, this value can be converted to characters as the integer value is the ascii value

// int getchar(void);


// int fgetc(FILE *stream);
// reads char from file, similar to getc it moves the fp to next char


// int ungetc(int ch, FILE *stream);
// this puts the character back into the input stream from where we just read

int main(int argc, char **argv) 
{
    char ch = '\0';
    //FILE *fp;

    //if (fp = fopen("file.txt", "rw")) {

        //ch = getc(fp);
        //ch = getc(stdin);
        //printf("char read : %c\n", ch);
        //while(ch != EOF) {
        //    ch = getc(fp);
        //}
      //  fclose(fp);
    //}
    //
    //

    // getc allows you to read from user file as well as stdin
    // whereas getchar only allows you to read from stdin

    //printf("%c\n", getchar());

    // always store the return type of getchar in an int not char type 


    /*
    int val = 0;
    while ((val = getchar()) != EOF) {
        printf("%c\n", val);
    }

    FILE *fp = NULL;

    char c = '\0';
    fp = fopen("file.txt", "r");
    if (fp == NULL) {
        // print error message
        exit(0);
    }

    while (1) {
        c = fgetc(fp);
        if (c == EOF) {
           // reached the end of file
            break;
        }
        // print the char c
    }
    fclose(fp);
    */


    // the below snippet reads all the spaces until a non space character is found
    // once a non space char is found it is put back into the c;
    char c = '\0';

    while (isspace(c = (char)getchar()));
    ungetc(c, stdin);

    printf("%c\n", c);
    return 0;
}
