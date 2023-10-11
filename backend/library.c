#include <stdio.h>
#include <stdlib.h>

int return_int() {
    return 30;
}

int* return_ints() {
    int* ints = malloc(sizeof(int) * 30);
    ints[0] = 30;
    ints[29] = 30;
    return ints;
}






