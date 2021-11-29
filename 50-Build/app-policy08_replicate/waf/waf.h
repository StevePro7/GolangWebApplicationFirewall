#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
//#include "modsecurity/transaction.h"
//#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
//RulesSet *rules = NULL;

int TheCAdd(int x, int y)
{
    modsec = msc_init();
    return x + y + 1;
}