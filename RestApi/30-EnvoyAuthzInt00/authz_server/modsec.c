#include <stdio.h>
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

void MyCInit()
{
    printf("MyCInit start\n");
    //printf("MyCInit fileName='%s'\n", file_name);
    if (modsec != NULL)
    {
        return;
    }

    modsec = msc_init();
    rules = msc_create_rules_set();

    const char *error = NULL;
    //msc_rules_add_file(rules, file_name, &error);
    msc_rules_add_file(rules, "rules/crs-setup.conf", &error);
    msc_rules_add_file(rules, "rules/modsecdefault.conf", &error);
    msc_rules_add_file(rules, "rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", &error);
    if (error != NULL)
    {
        printf("msc_rules_add_file error : '%s'\n", error);
    }

    //msc_rules_dump(rules);
    printf("MyCInit -end-\n");
    return;
}