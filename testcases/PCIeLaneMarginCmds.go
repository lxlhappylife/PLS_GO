package testcases

import (
	"PLS_GO/pcieutils"
	"fmt"
	"log"
	"time"
)

/*
- Recever Number
000b: Broadcast (Downstream Port Receiver and all Retimer Pseudo Port Receivers)
001b: Rx(A) (Downstream Port Receiver)
010b: Rx(B) (Retimer X or Z Upstream Pseudo Port Receiver)
30 011b: Rx(C) (Retimer X or Z Downstream Pseudo Port Receiver)
100b: Rx(D) (Retimer Y Upstream Pseudo Port Receiver)
101b: Rx(E) (Retimer Y Downstream Pseudo Port Receiver)
110b: Reserved

*/

//PCIeLaneMarginingHeaderAddress : the header address of PCIe Ext Cap 0x27 Lane Margining
// var PCIeLaneMarginingHeaderAddress uint16

// var bus uint8 = 0xa
// var device uint8 = 0x0
// var function uint8 = 0x0

// Device struct includes bus,device,function
type Device struct {
	bus      uint8
	device   uint8
	function uint8
}

// func init() {
// 	PCIeLaneMarginingHeaderAddress = pcieutils.GetPCIeExtCapHeaderAddress(bus,
// 		device,
// 		function,
// 		0x27)
// 	fmt.Printf("PCIeLaneMarginingHeaderAddress = 0x%x\n", PCIeLaneMarginingHeaderAddress)
// 	pcieutils.DisableHardwareAutonomousSpeed(bus, device, function) //EP
// 	pcieutils.DisableHardwareAutonomousWidth(bus, device, function) //EP
// 	pcieutils.DisableHardwareAutonomousSpeed(0x0, 0x3, 0x2)         //RC
// 	pcieutils.DisableHardwareAutonomousWidth(0x0, 0x3, 0x2)         //RC

// }

/*
PrePcieLanemarginCmds :
*/
// func PrePcieLanemarginCmds(EPBus uint8, EPDevice uint8, EPFunction uint8) {
// 	bus = EPBus
// 	device = EPDevice
// 	function = EPFunction
// 	PCIeLaneMarginingHeaderAddress = pcieutils.GetPCIeExtCapHeaderAddress(bus,
// 		device,
// 		function,
// 		0x27)
// }

// GetMarginingPortCapabilities : Get Margin use driver software
// return: 0x0 or 0x1
func GetMarginingPortCapabilities(bus uint8, device uint8, function uint8) uint {

	PCIeLaneMarginingHeaderAddress := pcieutils.GetPCIeExtCapHeaderAddress(bus,
		device,
		function,
		0x27)
	var offset uint16 = 0x4
	value := pcieutils.ConfigReadu16(bus, device, function, PCIeLaneMarginingHeaderAddress+offset)
	if value&0x1 == 0x1 {
		fmt.Println("Margining is partially implemented using Device Driver software.")
		return 0x1
	}
	fmt.Println("margining does not require device driver software.")
	return 0x0

}

//GetMarginingPortStatus : Get margining Port Status
// return: (MarginingReady, MarginingSoftwareReady)
func GetMarginingPortStatus(bus uint8, device uint8, function uint8) (uint, uint) {

	PCIeLaneMarginingHeaderAddress := pcieutils.GetPCIeExtCapHeaderAddress(bus,
		device,
		function,
		0x27)
	var offset uint16 = 0x6
	log.Printf("PCIeLaneMarginingHeaderAddress = 0x%x\n", PCIeLaneMarginingHeaderAddress)
	value := pcieutils.ConfigReadu8(bus, device, function, PCIeLaneMarginingHeaderAddress+offset)
	var MarginingReady, MarginingSoftwareReady uint
	MarginingReady = uint(value & 0x1)
	MarginingSoftwareReady = uint(value >> 4 & 0x1)
	log.Printf("value = 0x%x\n", value)
	if MarginingReady == 0x1 {
		fmt.Println("Margining Ready")
	} else {
		log.Fatal("Margining NOT Ready")
	}
	MarginingusesDriverSoftware := GetMarginingPortCapabilities(bus, device, function)
	if MarginingusesDriverSoftware == 0x1 {
		if MarginingSoftwareReady == 0x1 {
			fmt.Println("The required software has performed the required initialization.")
		} else {
			log.Fatal("The required software has not performed the required initialization.")
		}
	}
	return MarginingReady, MarginingSoftwareReady
}

/*
SetMarginingLaneControlRegister : Use for set Margining Lane Control Register
Input: LaneNum, ReceiverNum, MarinType, UsageModel, MarginPayload
Output: None
*/
func SetMarginingLaneControlRegister(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint, MarginType uint,
	UsageModel uint, MarginPayload uint8) {

	PCIeLaneMarginingHeaderAddress := pcieutils.GetPCIeExtCapHeaderAddress(bus,
		device,
		function,
		0x27)
	var value uint16 = 0x0
	var offset = uint16(0x8 + 0x4*LaneNum)
	value = uint16(ReceiverNum) | uint16(MarginType)<<3 | uint16(UsageModel)<<6 | uint16(MarginPayload)<<8
	pcieutils.ConfigWriteu16(bus, device, function, PCIeLaneMarginingHeaderAddress+offset, value)
}

/*
GetMarginingLaneStatusRegister : Use for get Margining Lane Status Register
Input: LaneNum
Output: ReceiverNum, MarinType, UsageModel, MarginPayload
*/
func GetMarginingLaneStatusRegister(bus uint8, device uint8, function uint8, LaneNum uint) (uint, uint, uint, uint8) {

	PCIeLaneMarginingHeaderAddress := pcieutils.GetPCIeExtCapHeaderAddress(bus,
		device,
		function,
		0x27)
	// var value uint16
	var offset = uint16(0xA + 0x4*LaneNum)
	value := pcieutils.ConfigReadu16(bus, device, function, PCIeLaneMarginingHeaderAddress+offset)
	ReceiverNum := (uint)(value & 0x7)
	MarginType := uint((value >> 3) & 0x7)
	UsageModel := uint((value >> 6) & 0x1)
	MarginPayload := uint8(value >> 8)
	return ReceiverNum, MarginType, UsageModel, MarginPayload
}

/*
NoCommand : Purpose of this step is to reset the Margining Lane Status config register before
issueing another command (which may have the same Margin Type encoding and Receiver Number.)
if 10ms expired since command was issued, declare Reciver failed margining and exit
*/
func NoCommand(bus uint8, device uint8, function uint8, LaneNum uint) {
	InputReceiverNum := uint(0x0)
	InputMarginType := uint(0x7)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x9C)
	log.Printf("Issue No Command\n")
	SetMarginingLaneControlRegister(bus, device, function,
		LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum, _, _, OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if 0x0 != OutputReceiverNum {
		log.Fatal("ReceiverNum Data should be ZERO")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
}

/*
ReportMarginControlCapabilites : Use for Read marin Control Capabilities
Input: LaneNum, ReceiverNum
Ouput:
*/
func ReportMarginControlCapabilites(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint) (bool, bool, bool, bool, bool) {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x88)

	log.Printf("Issue Report Margin Control Capabilites Command\n")
	SetMarginingLaneControlRegister(bus, device, function,
		LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum, OutputMarginType, OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	// log.Printf("InputReceiverNum = 0x%x\n", InputReceiverNum)
	// log.Printf("OutputReceiverNum = 0x%x\n", OutputReceiverNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	var IndErrorSampler bool
	var SampleReportingMethod bool
	var IndLeftRightTiming bool
	var IndUpDownVoltage bool
	var VoltageSupported bool
	if OutputMarginPayload&0x1 == 0x1 {
		IndErrorSampler = true
		log.Printf("IndErrorSampler Supports\n")
	} else {
		IndErrorSampler = false
	}
	if OutputMarginPayload>>1&0x1 == 0x1 {
		SampleReportingMethod = true
		log.Printf("SampleReportingMethod Supports\n")
	} else {
		SampleReportingMethod = false
	}
	if OutputMarginPayload>>2&0x1 == 0x1 {
		IndLeftRightTiming = true
		log.Printf("IndLeftRightTiming Supports\n")
	} else {
		IndLeftRightTiming = false
	}
	if OutputMarginPayload>>3&0x1 == 0x1 {
		IndUpDownVoltage = true
		log.Printf("IndUpDownVoltage Supports\n")
	} else {
		IndUpDownVoltage = false
	}
	if OutputMarginPayload>>4&0x1 == 0x1 {
		VoltageSupported = true
		log.Printf("VoltageSupported Supports\n")
	} else {
		VoltageSupported = false
	}
	return IndErrorSampler, SampleReportingMethod,
		IndLeftRightTiming, IndUpDownVoltage, VoltageSupported
}

/*
ReportNumVoltageSteps : ...
*/
func ReportNumVoltageSteps(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x89)
	log.Printf("Issue Report Num Voltage Steps Command.\n")
	SetMarginingLaneControlRegister(bus, device, function,
		LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x7F)
}

/*
ReportNumTimingSteps : ...
*/
func ReportNumTimingSteps(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8A)
	log.Printf("Issue Report Num Timing Steps Command.\n")
	SetMarginingLaneControlRegister(bus, device, function,
		LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x3F)
}

/*
ReportMaxTimingOffset : ...
*/
func ReportMaxTimingOffset(bus uint8, device uint8, function uint8,
	LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8B)
	log.Printf("Issue Report Max Timing Offset Command.\n")
	SetMarginingLaneControlRegister(bus, device, function,
		LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x7F)
}

/*
ReportMaxVoltageOffset : ...
*/
func ReportMaxVoltageOffset(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8C)
	log.Printf("Issue Report Max Voltage Offset Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x7F)
}

/*
ReportSamplingRateVoltage : ...
*/
func ReportSamplingRateVoltage(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8D)
	log.Printf("Issue Report Sampling Rate Voltage Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x3F)
}

/*
ReportSamplingRateTiming : ...
*/
func ReportSamplingRateTiming(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8E)
	log.Printf("Issue Report Sampling Rate Timing Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x3F)
}

/*
ReportSamepleCount : ...
*/
func ReportSamepleCount(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x8F)
	log.Printf("Issue Report Sampling Count Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x7F)
}

/*
ReportMaxLanes : ...
*/
func ReportMaxLanes(bus uint8, device uint8, function uint8, LaneNum uint, ReceiverNum uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x1)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x90)
	log.Printf("Issue Report Max Lanes Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x1F)
}

/*
SetErrorCountLimit : ...
*/
func SetErrorCountLimit(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint,
	ErrorCountLimit uint) uint8 {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x2) // Write Command
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(ErrorCountLimit) | (0x3 << 6)
	log.Printf("Issue Set Error Count Limit Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	log.Printf("OutputMarginPayload = 0x%x\n", OutputMarginPayload)
	return (OutputMarginPayload & 0x3F)
}

/*
GoToNormalSettings : ...
*/
func GoToNormalSettings(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint) {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x2) // Write Command
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0xF)
	log.Printf("Issue Go To Normal Settings Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	if InputMarginpayload != OutputMarginPayload {
		log.Fatal("Marginpayload Data Mis-Match")
	}
}

/*
ClearErrorLog : ...
*/
func ClearErrorLog(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint) {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x2) // Write Command
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := uint8(0x55)
	log.Printf("Issue Clear Error Log Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	if InputMarginpayload != OutputMarginPayload {
		log.Fatal("Marginpayload Data Mis-Match")
	}
}

/*
StepMarginToTimingOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/
func StepMarginToTimingOffset(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint, PayLoad uint8) (uint8, uint8) {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x3)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := PayLoad
	log.Printf("Issue Step Margin To Timing Offset Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	StepMarginExecutionStatus := OutputMarginPayload >> 6
	ErrorCount := OutputMarginPayload & 0x3F
	return StepMarginExecutionStatus, ErrorCount
}

/*
StepMarginToVoltageOffset :
Output: StepMarginExecutionStatus, ErrorCount
*/
func StepMarginToVoltageOffset(bus uint8, device uint8, function uint8, LaneNum uint,
	ReceiverNum uint, PayLoad uint8) (uint8, uint8) {
	InputReceiverNum := ReceiverNum
	InputMarginType := uint(0x4)
	InputUsageModel := uint(0x0) // Default
	InputMarginpayload := PayLoad
	log.Printf("Issue Step Margin To Voltage Offset Command.\n")
	SetMarginingLaneControlRegister(bus, device, function, LaneNum,
		InputReceiverNum,
		InputMarginType,
		InputUsageModel,
		InputMarginpayload)
	time.Sleep(10 * time.Millisecond)
	OutputReceiverNum,
		OutputMarginType,
		OutputUsageModel,
		OutputMarginPayload := GetMarginingLaneStatusRegister(bus, device, function, LaneNum)
	if InputReceiverNum != OutputReceiverNum {
		log.Fatal("ReceiverNum Data Mis-Match")
	}
	if InputMarginType != OutputMarginType {
		log.Fatal("MarginType Data Mis-Match")
	}
	if InputUsageModel != OutputUsageModel {
		log.Fatal("UsageModel Data Mis-Match")
	}
	StepMarginExecutionStatus := OutputMarginPayload >> 6
	ErrorCount := OutputMarginPayload & 0x3F
	return StepMarginExecutionStatus, ErrorCount
}
