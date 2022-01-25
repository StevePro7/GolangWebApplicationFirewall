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
    const char *error = NULL;
    fprintf(stderr, "C.LoadModSecurityCoreRuleSet() beg!!\n");
    msc_rules_add_file( rules, file_name, &error );

    if ( error != NULL )
    {
        fprintf(stderr, "bob '%s'\n", error);
    }
    fprintf(stderr, "C.LoadModSecurityCoreRuleSet() end??\n");
}