#include "waf.h"
#include <stdio.h>
#include <errno.h>

void Adriana()
{
    errno = 1;
    fprintf(stdout, "Hello from C..!!\n");
}