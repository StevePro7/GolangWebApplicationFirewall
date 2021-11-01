#include "main.h"

int main()
{
	char *array[] = { "one", "two", "three", "four" };
	//char *( *arrayPtr )[] = &array;

	printf( "Start!\n" );

	for( int i = 0; i < sizeof( array ) / sizeof( array[ 0 ] ); i++ ) 
	{
		//printf( "%s\n", ( *arrayPtr + i ) );
		printf( "%s\n", ( array[ i ] ) );
	}

	printf( "-end-!\n" );
	return 0;
}