#include "waf.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// Private helper function to initialize ModSecurity.
static void initializeModSecurityImpl();

void InitializeModSecurity()
{
    initializeModSecurityImpl();
}
static void initializeModSecurityImpl()
{
    modsec = msc_init();
//    msc_set_log_cb( modsec, null );
    rules = msc_create_rules_set();
}

void LoadModSecurityCoreRuleSet(char *file_name)
{
    fprintf(stderr, "C.LoadModSecurityCoreRuleSet()!!\n");
}