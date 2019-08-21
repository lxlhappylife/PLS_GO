package testcases

import (
	"PLS_GO/pcieutils"
	"log"
)

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
	log.Printf("LinkWidth = 0x%x\n", LinkWidth)
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
