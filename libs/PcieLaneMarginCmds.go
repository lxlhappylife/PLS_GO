package main

import (
	"C"
	// log "PLS_GO/log"
	"PLS_GO/testcases"
)

// GetMarginingPortCapabilities : Get Margin use driver software
// return : 0x0 or 0x1
//export GetMarginingPortCapabilities
func GetMarginingPortCapabilities(bus uint8, device uint8, function uint8) uint {
	return testcases.GetMarginingPortCapabilities(bus, device, function)
}

// GetMarginingReady : Get margining Port Status - MarginingReady
// return: (MarginingReady)
//export GetMarginingReady
func GetMarginingReady(bus uint8, device uint8, function uint8) uint {
	GetMarginingReady, _ := testcases.GetMarginingPortStatus(bus, device, function)
	return GetMarginingReady
}

// GetMarginingSoftwareReady : Get margining Port Status - MarginingSoftwareReady
// return: ( MarginingSoftwareReady)
//export GetMarginingSoftwareReady
func GetMarginingSoftwareReady(bus uint8, device uint8, function uint8) uint {
	_, MarginingSoftwareReady := testcases.GetMarginingPortStatus(bus, device, function)
	return MarginingSoftwareReady
}

/*
NoCommand : Purpose of this step is to reset the Margining Lane Status config register before
issueing another command (which may have the same Margin Type encoding and Receiver Number.)
if 10ms expired since command was issued, declare Reciver failed margining and exit
*/
//export NoCommand
func NoCommand(bus uint8, device uint8, function uint8, LaneNum uint) {
	testcases.NoCommand(bus, device, function, LaneNum)
}

// GetIndErrorSampler : Use for Read marin Control Capabilities - IndErrorSampler
//export GetIndErrorSampler
func GetIndErrorSampler(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) bool {
	IndErrorSampler, _, _, _, _ := testcases.ReportMarginControlCapabilites(bus, device, function, LaneNum,
		ReceiverNum)
	return IndErrorSampler
}

// GetSampleReportingMethod : Use for Read marin Control Capabilities - SampleReportingMethod
//export GetSampleReportingMethod
func GetSampleReportingMethod(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) bool {
	_, SampleReportingMethod, _, _, _ := testcases.ReportMarginControlCapabilites(bus, device, function, LaneNum,
		ReceiverNum)
	return SampleReportingMethod
}

// GetIndLeftRightTiming : Use for Read marin Control Capabilities - IndLeftRightTiming
//export GetIndLeftRightTiming
func GetIndLeftRightTiming(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) bool {
	_, _, IndLeftRightTiming, _, _ := testcases.ReportMarginControlCapabilites(bus, device, function, LaneNum,
		ReceiverNum)
	return IndLeftRightTiming
}

// GetIndUpDownVoltage : Use for Read marin Control Capabilities - IndUpDownVoltage
//export GetIndUpDownVoltage
func GetIndUpDownVoltage(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) bool {
	_, _, _, IndUpDownVoltage, _ := testcases.ReportMarginControlCapabilites(bus, device, function, LaneNum,
		ReceiverNum)
	return IndUpDownVoltage
}

// GetVoltageSupported : Use for Read marin Control Capabilities - VoltageSupported
//export GetVoltageSupported
func GetVoltageSupported(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) bool {
	_, _, _, _, VoltageSupported := testcases.ReportMarginControlCapabilites(bus, device, function, LaneNum,
		ReceiverNum)
	return VoltageSupported
}

/*
ReportNumVoltageSteps : ...
*/
//export ReportNumVoltageSteps
func ReportNumVoltageSteps(bus uint8, device uint8,
	function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportNumVoltageSteps(bus, device,
		function, LaneNum, ReceiverNum)
}

/*
ReportNumTimingSteps : ...
*/
//export ReportNumTimingSteps
func ReportNumTimingSteps(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportNumTimingSteps(bus, device, function,
		LaneNum, ReceiverNum)
}

/*
ReportMaxTimingOffset : ...
*/
//export ReportMaxTimingOffset
func ReportMaxTimingOffset(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportMaxTimingOffset(bus, device, function, LaneNum, ReceiverNum)
}

/*
ReportMaxVoltageOffset : ...
*/
//export ReportMaxVoltageOffset
func ReportMaxVoltageOffset(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportMaxVoltageOffset(bus, device, function,
		LaneNum, ReceiverNum)
}

/*
ReportSamplingRateVoltage : ...
*/
//export ReportSamplingRateVoltage
func ReportSamplingRateVoltage(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportSamplingRateVoltage(bus, device, function,
		LaneNum, ReceiverNum)
}

/*
ReportSamplingRateTiming : ...
*/
//export ReportSamplingRateTiming
func ReportSamplingRateTiming(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportSamplingRateTiming(bus, device, function,
		LaneNum, ReceiverNum)
}

/*
ReportSamepleCount : ...
*/
//export ReportSamepleCount
func ReportSamepleCount(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportSamepleCount(bus, device, function, LaneNum, ReceiverNum)
}

/*
ReportMaxLanes : ...
*/
//export ReportMaxLanes
func ReportMaxLanes(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	return testcases.ReportMaxLanes(bus, device, function, LaneNum, ReceiverNum)
}

/*
SetErrorCountLimit : ...
*/
//export SetErrorCountLimit
func SetErrorCountLimit(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint,
	ErrorCountLimit uint) uint8 {
	return SetErrorCountLimit(bus, device, function, LaneNum, ReceiverNum, ErrorCountLimit)
}

/*
GoToNormalSettings : ...
*/
//export GoToNormalSettings
func GoToNormalSettings(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint) {
	testcases.GoToNormalSettings(bus, device, function, LaneNum, ReceiverNum)
}

/*
ClearErrorLog : ...
*/
//export ClearErrorLog
func ClearErrorLog(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint) {
	testcases.ClearErrorLog(bus, device, function, LaneNum, ReceiverNum)
}

/*
StepMarginToTimingOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/
//export StepMarginToTimingOffset
func StepMarginToTimingOffset(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint, PayLoad uint8) (uint8, uint8) {
	return testcases.StepMarginToTimingOffset(bus, device, function,
		LaneNum, ReceiverNum, PayLoad)
}

/*
StepMarginToVoltageOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/
//export StepMarginToVoltageOffset
func StepMarginToVoltageOffset(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint, PayLoad uint8) (uint8, uint8) {
	return testcases.StepMarginToVoltageOffset(bus, device, function,
		LaneNum, ReceiverNum, PayLoad)
}

// func main() {}
