#include <stdio.h>
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

void MyCInit()
{
    printf("MyCInit start\n");
    if (modsec != NULL)
    {
        return;
    }

    modsec = msc_init();
    rules = msc_create_rules_set();

    const char *error = NULL;
    msc_rules_add_file(rules, "modsecdefault.conf", &error);
    msc_rules_add_file(rules, "crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);

    if (error != NULL)
    {
        printf("msc_rules_add_file error : '%s'\n", error);
    }

    printf("MyCInit -end-\n");
    return;
}

int MyCProcess(const char *uri)
{
    printf("MyCProcess start [%s]\n", uri);
    printf("MyCProcess -end- [%s]\n", uri);
    return 22;
}