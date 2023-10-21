#include <stdio.h>
int x = 1;
int y = 5;

int nwd(int x, int y) {
    while (x != 0 && y != 0) {
        if (x > y) {
            x = x % y;
        } else {
            y = y % x;
        }
    }
    printf("%d",x + y);
}

int main(void) {
    nwd(x,y);
    return 0;
}
