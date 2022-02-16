#include "waf.h"
//#include "_cgo_export.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

static void initializeModSecurityImpl()
{
    modsec = msc_init();
//    msc_set_log_cb(modsec, MyFuncPtr);  // TODO
    rules = msc_create_rules_set();
}

int AddRequestHeaders(char **reqHeaderKeys, char **reqHeaderVals, int size)
{
     int retVal = 0;

    Transaction *transaction = NULL;

    const char *reqHeaderKey;
    const char *reqHeaderVal;

    int index = 0;
    if (modsec == NULL)
    {
        initializeModSecurityImpl();
    }


    transaction = msc_new_transaction(modsec, rules, NULL);

    fprintf(stdout, "Beg\n");
    for( index = 0; index < size; index++ )
    {
        reqHeaderKey = reqHeaderKeys[ index ];
        reqHeaderVal = reqHeaderVals[ index ];

        fprintf(stdout, "Key : '%s'\n", reqHeaderKey);
        fprintf(stdout, "Val : '%s'\n", reqHeaderVal);

        retVal = msc_add_request_header(transaction, reqHeaderKey, reqHeaderVal);
        fprintf(stdout, "SGB : [%d]\n", retVal);
    }
    fprintf(stdout, "end\n");

    msc_transaction_cleanup(transaction);
    return 4;
}

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size)
{
    return calloc(sizeof(char*), size);
}
void freeCharArray(char **array, int size)
{
    int index;
    for (index = 0; index < size; index++)
    {
        free(array[index]);
    }

    free(array);
}
void setArrayString(char **array, char *input, int index)
{
    array[index] = input;
}
