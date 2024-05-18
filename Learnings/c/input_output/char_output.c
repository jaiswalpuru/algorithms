#include <stdio.h>

int main(int argc, char **argv)
{

    /*
    char ch = '\0';
    FILE *fp = NULL;
    
    if (fp = fopen("file.txt", "rw")) {
        ch = getc(fp);

        while (ch != EOF) {
            putc(ch, fp);// or we can use putc(ch, stdout);
            ch = getc(fp);
        }
        fclose(fp);
    }
    
    */
    char d = 'c';
    putc(d, stdout);

    // int puchar(int c);
    
    char *string = "Helo world!\n there we go";
    int i = 0;

    while (string[i] != '\0') {
        if (string[i] != '\n')
            putchar(string[i]);
        i++;
    }

    // int fputc(int char, FILE *fp);
    //

    FILE *fp = NULL;
    char c;

    fp = fopen("text.txt", "w");
    if (fp != NULL) {
        
        for (c = 'A'; c <= 'Z'; c++)
            fputc(c, fp);
        fclose(fp);
    }

    return 0;
}
