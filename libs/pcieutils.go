package main

import (
	"C"
	pcie "PLS_GO/pcieutils"
	"unsafe"
)

/*
GetHostBDF :
Output : RCBus,RCDevice,RCFunction uint8
*/
//export GetHostBDF
func GetHostBDF(bus uint8, device uint8,
	function uint8) uintptr {
	RC := make([]uint8, 3)
	RCBus, RCDevice, RCFunction := pcie.GetHostBDF(bus, device, function)
	RC = []uint8{RCBus, RCDevice, RCFunction}
	return uintptr(unsafe.Pointer(&RC[0]))
}

// func main() {}
