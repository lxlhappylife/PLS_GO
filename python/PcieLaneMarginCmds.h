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


// GetMarginingPortCapabilities : Get Margin use driver software
// return : 0x0 or 0x1

extern GoUint GetMarginingPortCapabilities(GoUint8 p0, GoUint8 p1, GoUint8 p2);

// GetMarginingReady : Get margining Port Status - MarginingReady
// return: (MarginingReady)

extern GoUint GetMarginingReady(GoUint8 p0, GoUint8 p1, GoUint8 p2);

// GetMarginingSoftwareReady : Get margining Port Status - MarginingSoftwareReady
// return: ( MarginingSoftwareReady)

extern GoUint GetMarginingSoftwareReady(GoUint8 p0, GoUint8 p1, GoUint8 p2);

/*
NoCommand : Purpose of this step is to reset the Margining Lane Status config register before
issueing another command (which may have the same Margin Type encoding and Receiver Number.)
if 10ms expired since command was issued, declare Reciver failed margining and exit
*/

extern void NoCommand(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3);

// GetIndErrorSampler : Use for Read marin Control Capabilities - IndErrorSampler

extern GoUint8 GetIndErrorSampler(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

// GetSampleReportingMethod : Use for Read marin Control Capabilities - SampleReportingMethod

extern GoUint8 GetSampleReportingMethod(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

// GetIndLeftRightTiming : Use for Read marin Control Capabilities - IndLeftRightTiming

extern GoUint8 GetIndLeftRightTiming(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

// GetIndUpDownVoltage : Use for Read marin Control Capabilities - IndUpDownVoltage

extern GoUint8 GetIndUpDownVoltage(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

// GetVoltageSupported : Use for Read marin Control Capabilities - VoltageSupported

extern GoUint8 GetVoltageSupported(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportNumVoltageSteps : ...
*/

extern GoUint8 ReportNumVoltageSteps(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportNumTimingSteps : ...
*/

extern GoUint8 ReportNumTimingSteps(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportMaxTimingOffset : ...
*/

extern GoUint8 ReportMaxTimingOffset(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportMaxVoltageOffset : ...
*/

extern GoUint8 ReportMaxVoltageOffset(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportSamplingRateVoltage : ...
*/

extern GoUint8 ReportSamplingRateVoltage(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportSamplingRateTiming : ...
*/

extern GoUint8 ReportSamplingRateTiming(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportSamepleCount : ...
*/

extern GoUint8 ReportSamepleCount(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ReportMaxLanes : ...
*/

extern GoUint8 ReportMaxLanes(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
SetErrorCountLimit : ...
*/

extern GoUint8 SetErrorCountLimit(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4, GoUint p5);

/*
GoToNormalSettings : ...
*/

extern void GoToNormalSettings(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/*
ClearErrorLog : ...
*/

extern void ClearErrorLog(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4);

/* Return type for StepMarginToTimingOffset */
struct StepMarginToTimingOffset_return {
	GoUint8 r0;
	GoUint8 r1;
};

/*
StepMarginToTimingOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/

extern struct StepMarginToTimingOffset_return StepMarginToTimingOffset(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4, GoUint8 p5);

/* Return type for StepMarginToVoltageOffset */
struct StepMarginToVoltageOffset_return {
	GoUint8 r0;
	GoUint8 r1;
};

/*
StepMarginToVoltageOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/

extern struct StepMarginToVoltageOffset_return StepMarginToVoltageOffset(GoUint8 p0, GoUint8 p1, GoUint8 p2, GoUint p3, GoUint p4, GoUint8 p5);

#ifdef __cplusplus
}
#endif
