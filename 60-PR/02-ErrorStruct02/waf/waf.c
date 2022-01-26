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

void LoadModSecurityCoreRuleSet( struct CoreRuleSetErrorObject *coreRuleSetErrorObject, char *file )
{
    const char *error = NULL;
    if ( modsec == NULL )
    {
        initializeModSecurityImpl();
    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        coreRuleSetErrorObject->error_message = (char *)error;
    }

    error = NULL;
}

/*
void LoadModSecurityCoreRuleSetX( struct ImgInfo *imgInfo, char *file )
{
    fprintf(stdout, "C.LoadModSecurityCoreRuleSet2() '%s'\n", file);

    const char *error = NULL;
    if ( modsec == NULL )
    {
        initializeModSecurityImpl();
    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        imgInfo->imgPath = (char *)error;
    }

    error = NULL;

    fprintf(stdout, "C.LoadModSecurityCoreRuleSet2() '%s'\n", file);
}
*/