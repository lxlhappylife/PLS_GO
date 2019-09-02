package testcases

import (
	"PLS_GO/pcieutils"
	// "PLS_GO/testcases"
	"fmt"

	log "PLS_GO/log"
)

const NAK uint8 = 0xFF
const TME uint8 = 0xFE // TME: Too Many Errors

func init() {
	// MarginingReady, MarginingSoftwareReady := GetMarginingPortStatus()
	// if GetMarginingPortCapabilities() == 0x1 {
	// 	if MarginingSoftwareReady == 0x0 {
	// 		log.Fatal("Device Driver Software is needed for PCIe Lane Margining Measurement.")
	// 	} else {
	// 		if MarginingReady == 0x0 {
	// 			log.Fatal("Lane Margining is not Ready")
	// 		}
	// 	}
	// } else {
	// 	if MarginingReady == 0x0 {
	// 		log.Fatal("Lane Margining is not Ready")
	// 	}
	// }
}

/*
PCIePrecondition :
*/
func PCIePrecondition(EPBus uint8, EPDevice uint8, EPFunction uint8,
	RCBus uint8, RCDevice uint8, RCFunction uint8) {
	// RCBus, RCDevice, RCFunction := pcieutils.GetHostBDF(EPBus, EPDevice, EPFunction)
	if pcieutils.CheckWhetherPCIeLinkSpeedReachGen4(EPBus, EPDevice, EPFunction) == false {
		log.Fatal("PCIe Lane Margining Feature only works when the link speed is 16GT/s.")
	}
	LinkWidth := pcieutils.GetPcieLinkWidth(EPBus, EPDevice, EPFunction)
	log.Info(fmt.Sprintf("LinkWidth = 0x%x\n", LinkWidth))
	// Disable ASPM
	pcieutils.DisableASPM(RCBus, RCDevice, RCFunction)
	pcieutils.DisableASPM(EPBus, EPDevice, EPFunction)
	// Disable Hardware Autonomous Speed
	pcieutils.DisableHardwareAutonomousSpeed(RCBus, RCDevice, RCFunction)
	pcieutils.DisableHardwareAutonomousSpeed(EPBus, EPDevice, EPFunction)
	// Disable Hardware Autonomous Width
	pcieutils.DisableHardwareAutonomousWidth(RCBus, RCDevice, RCFunction)
	pcieutils.DisableHardwareAutonomousWidth(EPBus, EPDevice, EPFunction)

}

/*
Precondition :
*/
func Precondition(EPBus uint8, EPDevice uint8, EPFunction uint8, LaneNum uint, ReceiverNum uint) {
	// ReadMarining Lane Status
	GetMarginingPortStatus(EPBus, EPDevice, EPFunction)
}

/*
DoPcieLaneMargining : Main Function for PCIe Lane manrgining
*/
func DoPcieLaneMargining(EPBus uint8, EPDevice uint8,
	EPFunction uint8, LaneNum uint, ReceiverNum uint) ([]uint8, []uint8, []uint8, []uint8) {
	RCBus, RCDevice, RCFunction := pcieutils.GetHostBDF(EPBus, EPDevice, EPFunction)
	PCIePrecondition(EPBus, EPDevice, EPFunction, RCBus, RCDevice, RCFunction)
	// Lane Num 0x0
	Precondition(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	var NumVoltageSteps, MaxVoltageOffset, SamplingRateVoltage uint8
	// IndErrorSampler,
	// 	SampleReportingMethod,
	// 	IndLeftRightTiming,
	// 	IndUpDownVoltage,
	_, _, IndLeftRightTiming,
		IndUpDownVoltage,
		VoltageSupported := ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	if VoltageSupported {
		// fmt.Printf("%s", "asdfasdf")
		NumVoltageSteps = ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
		log.PRINT(fmt.Sprintf("NumVoltageSteps = 0x%x\n", NumVoltageSteps), 2)
		NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		MaxVoltageOffset = ReportMaxVoltageOffset(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

		log.PRINT(fmt.Sprintf("MaxVoltageOffset = 0x%x\n", MaxVoltageOffset), 2)
		NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		SamplingRateVoltage = ReportSamplingRateVoltage(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

		log.PRINT(fmt.Sprintf("SamplingRateVoltage = 0x%x\n", SamplingRateVoltage), 2)
		NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	}
	NumTimingSteps := ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("NumTimingSteps = 0x%x\n", NumTimingSteps), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	MaxTimingOffset := ReportMaxTimingOffset(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("MaxTimingOffset = 0x%x\n", MaxTimingOffset), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	SamplingRateTiming := ReportSamplingRateTiming(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("SamplingRateTiming = 0x%x\n", SamplingRateTiming), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	SamepleCount := ReportSamepleCount(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("SamepleCount = 0x%x\n", SamepleCount), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	MaxLanes := ReportMaxLanes(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("MaxLanes = 0x%x\n", MaxLanes), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ErrorLimit := SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, 0x10) //Set Error Limit to 0x10

	log.PRINT(fmt.Sprintf("ErrorLimit = 0x%x\n", ErrorLimit), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	ClearErrorLog(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	// Do PCIe Lane Margin with Timing
	PayLoad := uint8(0x0)
	var TimingOffsetArray []uint8
	var LeftTimingOffsetArray []uint8
	var RightTimingOffsetArray []uint8
	for TimingOffset := uint8(0x0); TimingOffset < NumTimingSteps; TimingOffset++ {
		if IndLeftRightTiming == true {
			PayLoad = (0x1 << 6) // 0 : Right   1 : Left
			PayLoad |= TimingOffset
			fmt.Printf("TimingOffset=0x%x\n", TimingOffset)
			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("Left - TimingOffset = 0x%x - NAK\n", TimingOffset))

				LeftTimingOffsetArray = append(LeftTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("Left - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
			PayLoad = (0x0 << 6) // 0 : Right   1 : Left
			PayLoad |= TimingOffset

			StepMarginExecutionStatus, ErrorCount = StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("Right - TimingOffset = 0x%x - NAK\n", TimingOffset))
				RightTimingOffsetArray = append(RightTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("Right - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				RightTimingOffsetArray = append(RightTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				RightTimingOffsetArray = append(RightTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		} else {
			PayLoad = TimingOffset

			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - NAK\n", TimingOffset))
				TimingOffsetArray = append(TimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				TimingOffsetArray = append(TimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				TimingOffsetArray = append(TimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		}
	}

	// Do PCIe Lane Margin with Voltage
	PayLoad = uint8(0x0)
	var VoltageOffsetArray []uint8
	var UpVoltageOffsetArray []uint8
	var DownVoltageOffsetArray []uint8
	if VoltageSupported {
		for VoltageOffset := uint8(0x0); VoltageOffset < NumVoltageSteps; VoltageOffset++ {
			if IndUpDownVoltage == true {
				PayLoad = (0x1 << 7) // 0 : Up   1 : Down
				PayLoad |= VoltageOffset
				// fmt.Printf("VoltageOffset=0x%x\n", VoltageOffset)
				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("Down - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("Down - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
				PayLoad = (0x0 << 6) // 0 : Up   1 : Down
				PayLoad |= VoltageOffset

				StepMarginExecutionStatus, ErrorCount = StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("Up - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("Up - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			} else {
				PayLoad = VoltageOffset
				// fmt.Printf("VoltageOffset=0x%x\n", VoltageOffset)
				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					VoltageOffsetArray = append(VoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					VoltageOffsetArray = append(VoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					VoltageOffsetArray = append(VoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			}
		}
		if IndUpDownVoltage == true {
			log.Info(fmt.Sprintf("DownVoltageOffsetArray is %v\n", DownVoltageOffsetArray))
			log.Info(fmt.Sprintf("UpVoltageOffsetArray is %v\n", UpVoltageOffsetArray))
		} else {
			log.Info(fmt.Sprintf("VoltageOffsetArray is %v\n", VoltageOffsetArray))
			DownVoltageOffsetArray = VoltageOffsetArray
			UpVoltageOffsetArray = VoltageOffsetArray
		}
	} else {
		log.Warn("VoltageSupported does not support!\n")
	}
	if IndLeftRightTiming == true {
		log.Info(fmt.Sprintf("LeftTimingOffsetArray is %v\n", LeftTimingOffsetArray))
		log.Info(fmt.Sprintf("RightTimingOffsetArray is %v\n", RightTimingOffsetArray))
	} else {
		log.Info(fmt.Sprintf("TimingOffsetArray is %v\n", TimingOffsetArray))
		LeftTimingOffsetArray = TimingOffsetArray
		RightTimingOffsetArray = TimingOffsetArray
	}
	return DownVoltageOffsetArray, UpVoltageOffsetArray, LeftTimingOffsetArray, RightTimingOffsetArray
}

/*
GetPcieLaneMarginingLeftTimingOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
func GetPcieLaneMarginingLeftTimingOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) []uint8 {

	_, _, IndLeftRightTiming,
		_,
		_ := ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	NumTimingSteps := ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("NumTimingSteps = 0x%x\n", NumTimingSteps), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	ErrorLimit = SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, ErrorLimit)

	log.PRINT(fmt.Sprintf("ErrorLimit = 0x%x\n", ErrorLimit), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	ClearErrorLog(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	PayLoad := uint8(0x0)
	var LeftTimingOffsetArray []uint8
	for TimingOffset := uint8(0x0); TimingOffset < NumTimingSteps; TimingOffset++ {
		if IndLeftRightTiming == true {
			PayLoad = (0x1 << 6) // 0 : Right   1 : Left
			PayLoad |= TimingOffset
			fmt.Printf("TimingOffset=0x%x\n", TimingOffset)
			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("Left - TimingOffset = 0x%x - NAK\n", TimingOffset))

				LeftTimingOffsetArray = append(LeftTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("Left - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		} else {
			PayLoad = TimingOffset

			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - NAK\n", TimingOffset))
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				LeftTimingOffsetArray = append(LeftTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		}

	}
	return LeftTimingOffsetArray
}

/*
GetPcieLaneMarginingRightTimingOffset : Main Function for PCIe Lane manrgining LeftTimingOffsetArray
*/
func GetPcieLaneMarginingRightTimingOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) []uint8 {

	_, _, IndLeftRightTiming,
		_,
		_ := ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	NumTimingSteps := ReportNumTimingSteps(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum)

	log.PRINT(fmt.Sprintf("NumTimingSteps = 0x%x\n", NumTimingSteps), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	ErrorLimit = SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, ErrorLimit)

	log.PRINT(fmt.Sprintf("ErrorLimit = 0x%x\n", ErrorLimit), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	ClearErrorLog(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	PayLoad := uint8(0x0)
	var RightTimingOffsetArray []uint8
	for TimingOffset := uint8(0x0); TimingOffset < NumTimingSteps; TimingOffset++ {
		if IndLeftRightTiming == true {
			PayLoad = (0x1 << 6) // 0 : Right   1 : Left
			PayLoad |= TimingOffset
			fmt.Printf("TimingOffset=0x%x\n", TimingOffset)
			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("Right - TimingOffset = 0x%x - NAK\n", TimingOffset))

				RightTimingOffsetArray = append(RightTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("Right - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				RightTimingOffsetArray = append(RightTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				RightTimingOffsetArray = append(RightTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		} else {
			PayLoad = TimingOffset

			StepMarginExecutionStatus, ErrorCount := StepMarginToTimingOffset(EPBus, EPDevice, EPFunction,
				LaneNum, ReceiverNum, PayLoad)
			switch StepMarginExecutionStatus {
			case 0x3: //NAK
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - NAK\n", TimingOffset))
				RightTimingOffsetArray = append(RightTimingOffsetArray, NAK)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x0: // TOO MANY ERROR
				log.Error(fmt.Sprintf("BothSides - TimingOffset = 0x%x - TOO MANY ERROR\n", TimingOffset))
				RightTimingOffsetArray = append(RightTimingOffsetArray, TME)
				NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
			case 0x2:
				RightTimingOffsetArray = append(RightTimingOffsetArray, ErrorCount)
			case 0x1:
				log.Fatal("StepMarginExecutionStatus = 0x1\n")
			default:
				log.Fatal("unsupport StepMarginExecutionStatus type\n")
			}
		}

	}
	return RightTimingOffsetArray
}

/*
GetPcieLaneMarginingUpVoltageOffset : Main Function for PCIe Lane manrgining UpVoltageOffset
*/
func GetPcieLaneMarginingUpVoltageOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) []uint8 {
	var NumVoltageSteps uint8
	_, _, _, IndUpDownVoltage,
		VoltageSupported := ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ErrorLimit = SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, ErrorLimit)
	log.PRINT(fmt.Sprintf("ErrorLimit = 0x%x\n", ErrorLimit), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ClearErrorLog(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	PayLoad := uint8(0x0)
	var UpVoltageOffsetArray []uint8
	if VoltageSupported {
		NumVoltageSteps = ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
		log.PRINT(fmt.Sprintf("NumVoltageSteps = 0x%x\n", NumVoltageSteps), 2)
		NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		for VoltageOffset := uint8(0x0); VoltageOffset < NumVoltageSteps; VoltageOffset++ {
			if IndUpDownVoltage == true {
				PayLoad = (0x0 << 6) // 0 : Up   1 : Down
				PayLoad |= VoltageOffset

				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("Up - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("Up - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			} else {
				PayLoad = VoltageOffset
				// fmt.Printf("VoltageOffset=0x%x\n", VoltageOffset)
				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					UpVoltageOffsetArray = append(UpVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			}
		}
	} else {
		log.Warn("VoltageSupported does not support!\n")
		UpVoltageOffsetArray = []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	}
	return UpVoltageOffsetArray
}

/*
GetPcieLaneMarginingDownVoltageOffset : Main Function for PCIe Lane manrgining DownVoltageOffset
*/
func GetPcieLaneMarginingDownVoltageOffset(EPBus uint8,
	EPDevice uint8,
	EPFunction uint8,
	LaneNum uint,
	ReceiverNum uint,
	ErrorLimit uint) []uint8 {
	var NumVoltageSteps uint8
	_, _, _, IndUpDownVoltage,
		VoltageSupported := ReportMarginControlCapabilites(EPBus,
		EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ErrorLimit = SetErrorCountLimit(EPBus, EPDevice, EPFunction,
		LaneNum, ReceiverNum, ErrorLimit)
	log.PRINT(fmt.Sprintf("ErrorLimit = 0x%x\n", ErrorLimit), 2)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
	ClearErrorLog(EPBus, EPDevice, EPFunction, LaneNum, ReceiverNum)
	NoCommand(EPBus, EPDevice, EPFunction, LaneNum)

	PayLoad := uint8(0x0)
	var DownVoltageOffsetArray []uint8
	if VoltageSupported {
		NumVoltageSteps = ReportNumVoltageSteps(EPBus, EPDevice, EPFunction,
			LaneNum, ReceiverNum)
		log.PRINT(fmt.Sprintf("NumVoltageSteps = 0x%x\n", NumVoltageSteps), 2)
		NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
		for VoltageOffset := uint8(0x0); VoltageOffset < NumVoltageSteps; VoltageOffset++ {
			if IndUpDownVoltage == true {
				PayLoad = (0x1 << 6) // 0 : Up   1 : Down
				PayLoad |= VoltageOffset

				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("Down - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("Down - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			} else {
				PayLoad = VoltageOffset
				// fmt.Printf("VoltageOffset=0x%x\n", VoltageOffset)
				StepMarginExecutionStatus, ErrorCount := StepMarginToVoltageOffset(EPBus, EPDevice, EPFunction,
					LaneNum, ReceiverNum, PayLoad)
				switch StepMarginExecutionStatus {
				case 0x3: //NAK
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - NAK\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, NAK)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x0: // TOO MANY ERROR
					log.Error(fmt.Sprintf("BothSides - VoltageOffset = 0x%x - TOO MANY ERROR\n", VoltageOffset))
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, TME)
					NoCommand(EPBus, EPDevice, EPFunction, LaneNum)
				case 0x2:
					DownVoltageOffsetArray = append(DownVoltageOffsetArray, ErrorCount)
				case 0x1:
					log.Fatal("StepMarginExecutionStatus = 0x1\n")
				default:
					log.Fatal("unsupport StepMarginExecutionStatus type\n")
				}
			}
		}
	} else {
		log.Warn("VoltageSupported does not support!\n")
		DownVoltageOffsetArray = []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	}
	return DownVoltageOffsetArray
}

/*
MeasurePCIeLaneMariningFromBothSide :
*/
func MeasurePCIeLaneMariningFromBothSide(EPBus uint8, EPDevice uint8, EPFunction uint8, Population int) {
	RCBus, RCDevice, RCFunction := pcieutils.GetHostBDF(EPBus, EPDevice, EPFunction)
	PCIePrecondition(EPBus, EPDevice, EPFunction,
		RCBus, RCDevice, RCFunction)
	/*
		Root Port or Switch
		|				RxA(001b)
		|				^
		V				|
		RxB(010b)		|
		####Retimer 1####
		|				RxC(011b)
		|				^
		V				|
		RxD(100b)		|
		####Retimer 2####
		|				RxE(101b)
		|				^
		V				|
		RxF(110b)		|
		ENDPOINT or Switch
	*/
	RCNumberOfRetimers := pcieutils.GetNumOfRetimer(RCBus, RCDevice, RCFunction)
	log.PRINTF(fmt.Sprintf("RCNumberOfRetimers=%x\n", RCNumberOfRetimers), 2, "Test.log")
}
