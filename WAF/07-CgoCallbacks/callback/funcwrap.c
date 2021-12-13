#include <stdio.h>
#include "_cgo_export.h"
#include "calladd.h"

void pass_GoAdd(void)
{
    printf("C  pass_GoAdd beg\n");
    call_add(&GoAdd);
    printf("C  pass_GoAdd end\n");
}