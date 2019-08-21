package main

import (
	"PLS_GO/pcieutils"
	"PLS_GO/testcases"
	"log"
	// "fmt"
)

func main() {
	// dut = Device
	// testcases.PrePcieLanemarginCmds()
	EPBus := uint8(0xa)
	EPDevice := uint8(0x0)
	EPFunction := uint8(0x0)
	RCBus, RCDevice, RCFunction := pcieutils.GetHostBDF(EPBus, EPDevice, EPFunction)
	testcases.PCIePrecondition(EPBus, EPDevice, EPFunction, RCBus, RCDevice, RCFunction)
	// Lane Num 0x0
	LaneNum := uint(0x0)
	ReceiverNum := uint(0x6)
	testcases.Precondition(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

	// IndErrorSampler,
	// 	SampleReportingMethod,
	// 	IndLeftRightTiming,
	// 	IndUpDownVoltage,
	_, _, _, _, VoltageSupported := testcases.ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	if VoltageSupported {
		NumVoltageSteps := testcases.ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
		log.Printf("NumVoltageSteps = 0x%x\n", NumVoltageSteps)
		testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		MaxVoltageOffset := testcases.ReportMaxVoltageOffset(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

		log.Printf("MaxVoltageOffset = 0x%x\n", MaxVoltageOffset)
		testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		SamplingRateVoltage := testcases.ReportSamplingRateVoltage(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

		log.Printf("SamplingRateVoltage = 0x%x\n", SamplingRateVoltage)
		testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	}
	NumTimingSteps := testcases.ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.Printf("NumTimingSteps = 0x%x\n", NumTimingSteps)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	MaxTimingOffset := testcases.ReportMaxTimingOffset(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

	log.Printf("MaxTimingOffset = 0x%x\n", MaxTimingOffset)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	SamplingRateTiming := testcases.ReportSamplingRateTiming(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.Printf("SamplingRateTiming = 0x%x\n", SamplingRateTiming)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	SamepleCount := testcases.ReportSamepleCount(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.Printf("SamepleCount = 0x%x\n", SamepleCount)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	MaxLanes := testcases.ReportMaxLanes(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.Printf("MaxLanes = 0x%x\n", MaxLanes)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ErrorLimit := testcases.SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, 0x5) //Set Error Limit to 0x5

	log.Printf("ErrorLimit = 0x%x\n", ErrorLimit)
	testcases.NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
}
