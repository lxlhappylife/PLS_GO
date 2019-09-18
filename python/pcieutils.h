/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

#endif

/* Start of preamble from import "C" comments.  */




/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef _GoString_ GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


/*
GetHostBDF :
Output : RCBus,RCDevice,RCFunction uint8
*/

extern GoUintptr GetHostBDF(GoUint8 p0, GoUint8 p1, GoUint8 p2);

// ConfigReadu32 is for read device's Config Space and return 32 bits value

extern GoUint32 ConfigReadu32(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3);

// ConfigReadu16 is for read device's Config Space and return 16 bits value

extern GoUint16 ConfigReadu16(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3);

// ConfigReadu8 is for read device's Config Space and return 8 bits value

extern GoUint8 ConfigReadu8(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3);

// ConfigWriteu32 is for Write device's Config Space with 32 bits value

extern void ConfigWriteu32(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3, GoUint32 p4);

// ConfigWriteu16 is for Write device's Config Space with 16 bits value

extern void ConfigWriteu16(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3, GoUint16 p4);

// ConfigWriteu8 is for Write device's Config Space with 8 bits value

extern void ConfigWriteu8(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint16 p3, GoUint8 p4);

#ifdef __cplusplus
}
#endif
