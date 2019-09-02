//LaneMargin.go
package main

import (
	"C"
	log "PLS_GO/log"
	"PLS_GO/testcases"
	"fmt"
	"unsafe"
)

/*
GetLaneMargin : Get the four arrays
- LeftTiming,
- RightTiming,
- UpVoltage,	(Optional)
- DownVoltage,	(Optional)
*/
//export GetLaneMargin
func GetLaneMargin(bus uint8, device uint8,
	function uint8, LaneNum uint,
	ReceiverNum uint) {

	fmt.Printf("bus = %d\ndevice = %d\nfunction = %d\n", bus,
		device, function)
	log.Info(fmt.Sprintf("bus = %x\ndevice = %x\nfunction = %x\n", uint8(bus),
		uint8(device), uint8(function)))
	// log.Fatal("0\n")
	DownVoltageOffsetArray,
		UpVoltageOffsetArray,
		LeftTimingOffsetArray,
		RightTimingOffsetArray := testcases.DoPcieLaneMargining(bus,
		device, function, LaneNum, ReceiverNum)
	fmt.Printf("%v", DownVoltageOffsetArray)
	fmt.Printf("%v", UpVoltageOffsetArray)
	fmt.Printf("%v", LeftTimingOffsetArray)
	fmt.Printf("%v", RightTimingOffsetArray)
	// return LeftTimingOffsetArray,
	// 	RightTimingOffsetArray,
	// 	UpVoltageOffsetArray,
	// 	DownVoltageOffsetArray
}

/*
GetPcieLaneMarginingLeftTimingOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
//export GetPcieLaneMarginingLeftTimingOffset
func GetPcieLaneMarginingLeftTimingOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) uintptr {
	NumTimingSteps := testcases.ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)
	LeftTimingOffsetArray := make([]uint8, int(NumTimingSteps))
	LeftTimingOffsetArray = testcases.GetPcieLaneMarginingLeftTimingOffset(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum, ErrorLimit)
	return uintptr(unsafe.Pointer(&LeftTimingOffsetArray[0]))
}

/*
GetPcieLaneMarginingRightTimingOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
//export GetPcieLaneMarginingRightTimingOffset
func GetPcieLaneMarginingRightTimingOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) uintptr {
	NumTimingSteps := testcases.ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)
	RightTimingOffsetArray := make([]uint8, int(NumTimingSteps))
	RightTimingOffsetArray = testcases.GetPcieLaneMarginingLeftTimingOffset(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum, ErrorLimit)
	return uintptr(unsafe.Pointer(&RightTimingOffsetArray[0]))
}

/*
GetPcieLaneMarginingUpVoltageOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
//export GetPcieLaneMarginingUpVoltageOffset
func GetPcieLaneMarginingUpVoltageOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) uintptr {
	_, _, _, _,
		VoltageSupported := testcases.ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	var NumTimingSteps uint8
	if VoltageSupported {
		NumTimingSteps = testcases.ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
	} else {
		NumTimingSteps = 0x6

	}
	UpVoltageOffsetArray := make([]uint8, int(NumTimingSteps))
	UpVoltageOffsetArray = testcases.GetPcieLaneMarginingUpVoltageOffset(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum, ErrorLimit)

	return uintptr(unsafe.Pointer(&UpVoltageOffsetArray[0]))
}

/*
GetPcieLaneMarginingDownVoltageOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
//export GetPcieLaneMarginingDownVoltageOffset
func GetPcieLaneMarginingDownVoltageOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) uintptr {
	_, _, _, _,
		VoltageSupported := testcases.ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	var NumTimingSteps uint8
	if VoltageSupported {
		NumTimingSteps = testcases.ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
	} else {
		NumTimingSteps = 0x6

	}
	DownVoltageOffsetArray := make([]uint8, int(NumTimingSteps))
	DownVoltageOffsetArray = testcases.GetPcieLaneMarginingDownVoltageOffset(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum, ErrorLimit)

	return uintptr(unsafe.Pointer(&DownVoltageOffsetArray[0]))
}
func main() {}
