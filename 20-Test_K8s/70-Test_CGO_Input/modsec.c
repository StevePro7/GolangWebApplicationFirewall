#include <stdio.h>

void MyCInit()
{
    printf("MyCInit start\n");
    /*
    const char *error = NULL;
    msc_rules_add_file(rules, "modsecdefault.conf", &error);
    msc_rules_add_file(rules, "crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
    if (error != NULL)
    {
        printf("msc_rules_add_file error : '%s'\n", error);
    }
    */
    printf("MyCInit -end-\n");
    return;
}
