#include <stdio.h>
#include "calladd.h"

void
call_add(adder func) {
    printf("Calling adder\n");
    int i = func(3, 4);
    printf("Adder returned %d\n", i);
}