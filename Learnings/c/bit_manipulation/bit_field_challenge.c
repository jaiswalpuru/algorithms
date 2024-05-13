#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define SOLID 0
#define DOTTED 1
#define DASHED 2

#define BLUE 4
#define GREEN 2
#define RED 1

#define BLACK 0
#define YELLOW (RED | GREEN)
#define MAGENTA (RED | BLUE)
#define CYAN (GREEN | BLUE)
#define WHITE (RED | GREEN | BLUE) 

const char * colors[8] = {"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"};

struct box_props {
    bool opaque                 : 1;
    unsigned int fill_color     : 3;
    unsigned int                : 4;//padding
    bool show_border            : 1;
    unsigned int border_color   : 3;
    unsigned int border_style   : 2;
    unsigned int                : 2;
};

void show_settings(const struct box_props *bp);

int main(int argc, char **argv)
{
    struct box_props props = {true, YELLOW, true, GREEN, DASHED};
    printf("Original box settings\n");
    show_settings(&props);
    return 0;   
}

void show_settings(const struct box_props *bp)
{
    printf("box is %s.\n", bp->opaque == true ? "opaque":"transparent");
    printf("fill color %s\n", colors[bp->fill_color]);
}

