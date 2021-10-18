#include <stdio.h>
#include "modsecurity/modsecurity.h"

ModSecurity *modsec = NULL;

void MyCInit()
{
    printf("MyCInit start!\n");
    if (modsec != NULL)
    {
        return;
    }

    printf("msc_init() start\n");
    modsec = msc_init();
    printf("msc_init() -end-\n");

    printf("MyCInit -end-!\n");
}