#include <stdio.h>
//#include "_cgo_export.h"

void InitializeModSecurity()
{
    fprintf(stderr, "C InitializeModSecurity beg\n");
    fprintf(stderr, "C InitializeModSecurity end\n");
}

void pass_GoAdd(void)
{
    fprintf(stderr, "C pass_GoAdd beg\n");
    call_add(&GoAdd);
    fprintf(stderr, "C pass_GoAdd end\n");
}