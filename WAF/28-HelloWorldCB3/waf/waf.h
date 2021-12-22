// Important: <stdlib.h> header file required for Golang statements e.g. defer C.free(unsafe.Pointer(API))
#include <stdlib.h>

// ModSecurity logging callback infrastructure APIs.
typedef void( *ModSecurityLoggingCallbackFunctionPointer )( char * );
extern void InvokeModSecurityLoggingCallback( ModSecurityLoggingCallbackFunctionPointer, char * );

void InitializeModSecurity();
int LoadModSecurityCoreRuleSet( char **array, int size );
int ProcessHttpRequest( char *id, char *uri, char *http_method, char *http_protocol, char *http_version, char *client_host, int client_port, char *server_host, int server_port );
void CleanupModSecurity();

// Helper functions to store all core rule set file names in memory.
char **makeCharArray( int size );
void freeCharArray( char **array, int size );
void setArrayString( char **array, char *filename, int index );
