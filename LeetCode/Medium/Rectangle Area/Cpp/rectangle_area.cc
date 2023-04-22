#include <iostream>

int area(int x1, int y1, int x2, int y2) { return (x2-x1)*(y2-y1); }

int rectangle_area(int ax1, int ay1, int ax2, int ay2, int bx1, int by1, int bx2, int by2) {
    int x1 = std::max(ax1, bx1);
    int y1 = std::max(ay1, by1);
    int x2 = std::min(ax2, bx2);
    int y2 = std::min(ay2, by2);
    return (x2 > x1 && y2 > y1)?area(ax1, ay1, ax2, ay2) + area(bx1, by1, bx2, by2) - area(x1, y1, x2, y2):area(ax1, ay1, ax2, ay2) + area(bx1, by1, bx2, by2);
}

int main() {
    std::cout<<rectangle_area(-2, -2, 2, 2, -2, -2, 2, 2)<<"\n";
}
