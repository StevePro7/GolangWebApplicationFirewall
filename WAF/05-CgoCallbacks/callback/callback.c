#include <stdio.h>
#include "calladd.h"

void call_add(adder func) {
    printf("C  Callback beg\n");
    int i = func(10, 20);
    printf("C  Callback end [%d]\n", i);
}