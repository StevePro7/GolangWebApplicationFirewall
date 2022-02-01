#include "waf.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"
#include <stdio.h>

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// General public APIs.
void InitializeModSecurity()
{
    modsec = msc_init();
    //msc_set_log_cb( modsec, CModSecurityLoggingCallback );
    rules = msc_create_rules_set();
}

const char* LoadModSecurityCoreRuleSet( char *file )
{
    const char *error = NULL;
    char *error_message = NULL;
//    if ( modsec == NULL )
//    {
//        initializeModSecurityImpl();
//    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        error_message = (char *)error;
    }

    error = NULL;
    return error_message;
}

void CleanupModSecurity()
{
    if ( rules != NULL )
    {
        msc_rules_cleanup( rules );
    }
    if ( modsec != NULL )
    {
        msc_cleanup( modsec );
    }

    rules = NULL;
    modsec = NULL;
}

int ProcessHttpRequest( char *id, char *uri, char *http_method, char *http_protocol, char *http_version, char *client_host, int client_port, char *server_host, int server_port )
{
    int detection = 0;
    int test1 = 0;
    int test2 = 0;
    int test3 = 0;
    int test4 = 0;

    Transaction *transaction = NULL;
    transaction = msc_new_transaction_with_id( modsec, rules, id, NULL );

    test1 = msc_process_connection( transaction, client_host, client_port, server_host, server_port );
    fprintf(stdout, "test1 = %d\n", test1);
    test1 = 0;
    //fprintf(stdout, "test1 = %d\n", test1);
    if ( !test1 )
    {
        //fprintf(stdout, "error1!!\n");
    }

    test2 = msc_process_uri( transaction, uri, http_protocol, http_version );
    fprintf(stdout, "test2 = %d\n", test2);
    test2 = 0;
    //fprintf(stdout, "test2 = %d\n", test2);
    if ( !test2 )
    {
        //fprintf(stdout, "error2!!\n");
    }

    test3 = msc_process_request_headers( transaction );
    fprintf(stdout, "test3 = %d\n", test3);
    test3 = 0;
    //fprintf(stdout, "test3 = %d\n", test3);
    if ( !test3 )
    {
        //fprintf(stdout, "error3!!\n");
    }

    test4 = msc_process_request_body( transaction );
    fprintf(stdout, "test4 = %d\n", test4);
    test4 = 0;
    //fprintf(stdout, "test4 = %d\n", test4);
    if ( !test4 )
    {
        //fprintf(stdout, "error2!!\n");
    }

        ModSecurityIntervention intervention;
        intervention.status = 200;
        intervention.url = NULL;
        intervention.log = NULL;
        intervention.disruptive = 0;

        detection = msc_intervention( transaction, &intervention );
        fprintf(stdout, "detection = %d\n", detection);
        if ( transaction != NULL )
        {
            msc_transaction_cleanup( transaction );
            transaction = NULL;
        }

        return detection;
}