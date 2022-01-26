// Important: <stdlib.h> header file required for Golang statements e.g. defer C.free(unsafe.Pointer(API))
#include <stdlib.h>

// ModSecurity logging callback infrastructure APIs.
typedef void( *ModSecurityLoggingCallbackFunctionPointer )( char * );
extern void InvokeModSecurityLoggingCallback( ModSecurityLoggingCallbackFunctionPointer, char * );

void InitializeModSecurity();
const char * LoadModSecurityCoreRuleSet2( char **array, int size);
void CleanupModSecurity();

// Helper functions to store all core rule set file names in memory.
char **makeCharArray( int size );
void freeCharArray( char **array, int size );
void setArrayString( char **array, char *filename, int index );