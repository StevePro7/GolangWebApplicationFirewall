#include <stdio.h>
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

void MyCInit()
{
    printf("!MyCInit start!\n");
    if (modsec != NULL)
    {
        return;
    }

    //printf("msc_init() start\n");
    modsec = msc_init();
    //printf("msc_init() -end-\n");

    //printf("msc_create_rules_set() start\n");
    rules = msc_create_rules_set();
    //printf("msc_create_rules_set() -end-\n");

    const char *error = NULL;
    printf("msc_rules_add_file() start\n");
    msc_rules_add_file(rules, "modsecdefault.conf", &error);
    msc_rules_add_file(rules, "crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
    if (error != NULL)
    {
        printf("The Error '%s'\n", error);
    }
    //printf("msc_create_rules_set() start\n");
    //rules = msc_create_rules_set();
    //printf("msc_create_rules_set() -end-\n");

    printf("msc_rules_add_file() -end-\n");
    printf("!!MyCInit -end-!\n");
    return;
}