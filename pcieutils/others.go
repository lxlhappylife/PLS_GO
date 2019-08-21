// others.go File

// This file was mostly used for pcie access related features
// -- Chi Zhang 7/10/2019

//Package pcieutils provides some api for Config/Mem space read access.
package pcieutils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var counterForOneMicrosecond int64

// ERRORu8 : cannot find matched capID
var ERRORu8 uint8 = 0xFE

// ERRORu16 : cannot find matched Ext capID
var ERRORu16 uint16 = 0xFFFE

func delayUnit() {
	for i := 0; i < 100; i++ {
		const a = 0
	}
}
func init() {
	// var count int64
	// running := true
	// go func() {
	// 	for running {
	// 		delayUnit()
	// 		count++
	// 	}
	// 	log.Println("go Done.")
	// }()
	// time.Sleep(time.Second)
	// running = false
	// counterForOneMicrosecond = int64(count / (1000 * 1000))
	// log.Printf("counterForOneMicrosecond = %d\n", counterForOneMicrosecond)
}

// Polling is continuely issue config/Mem Read to a special address
func Polling(baseAddress int64, offset uint16, delayUs int64, durationSec int64) uint32 {
	func() {
		var count int64
		running := true
		go func() {
			for running {
				delayUnit()
				count++
			}
			log.Println("go Done.")
		}()
		time.Sleep(time.Second)
		running = false
		counterForOneMicrosecond = int64(count / (1000 * 1000))
		log.Printf("counterForOneMicrosecond = %d\n", counterForOneMicrosecond)
	}()
	var value uint32 = 0xFFFFFFFF
	const bufferSize int = 4096

	file, err := os.Open("/dev/mem")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	mmap, err := syscall.Mmap(int(file.Fd()), baseAddress, bufferSize, syscall.PROT_READ, syscall.MAP_SHARED)

	if err != nil {
		log.Fatal(err)
	}
	// ticker := time.NewTicker(time.Duration(delayUs) * time.Millisecond)
	// go func() {
	// 	for range ticker.C {
	// 		value = *(*uint32)(unsafe.Pointer(&mmap[offset]))
	// 	}
	// }()
	// time.Sleep(time.Duration(durationSec) * time.Second)
	// ticker.Stop()
	start := time.Now()
	var count int64
	for time.Duration(durationSec)*time.Second >= time.Since(start) {
		count = 0
		value = *(*uint32)(unsafe.Pointer(&mmap[offset]))

		for count < delayUs*counterForOneMicrosecond {
			delayUnit()
			count++
		}
	}
	err = syscall.Munmap(mmap)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// GetPCIeBaseAddress : Read PCIe base address
func GetPCIeBaseAddress() uint64 {
	var address uint64
	// var cmd = "cat /proc/iomem | grep 'PCI MMCONFIG 000'"
	// var cmd = "cat /proc/iomem"
	buf, err := exec.Command("cat", "/proc/iomem").Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(buf, []byte("\n"))
	var targetLine, addressString string
	for num := 0; num < len(lines); num++ {
		// line_str = string(lines[num])
		if strings.Contains(string(lines[num]), "PCI MMCONFIG 000") {
			targetLine = string(lines[num])
		}
	}
	// fmt.Printf("%s ,", targetLine)
	if targetLine != "" {
		items := strings.Split(targetLine, "-")
		addressString = items[0]
	} else {
		log.Fatal("Cannot find target line in iomem.")
		return 0xFFFFFFFF
	}
	address, err = strconv.ParseUint(addressString, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return address
}

// GetPCIeCapHeaderAddress : calcuate the header address
func GetPCIeCapHeaderAddress(bus uint8, device uint8, function uint8, CapID uint8) uint8 {
	var startAddress uint8 = 0x34
	var address = ConfigReadu8(bus, device, function, uint16(startAddress))
	var value uint8
	var count int

	fmt.Printf("%x-%x-%x\n", bus, device, function)
	for true {
		value = ConfigReadu8(bus, device, function, uint16(address))
		if value == CapID {
			return address
		}
		address = ConfigReadu8(bus, device, function, uint16(address+0x1))
		if address == 0x0 {
			log.Fatal("Scanning to the end.")
			return ERRORu8
		}
		if count == 100 {
			log.Fatal("Cannot find match CapID")
			break
		} else {
			count++
		}
	}
	return ERRORu8
}

// GetPCIeExtCapHeaderAddress : calcuate the header address of PCIe Ext Cap
/*Bit Location Register Description Attributes
15:0 PCI Express Extended Capability ID  This field is a PCISIG defined ID number that indicates the nature and format
of the Extended Capability.
The Extended Capability ID for the Physical Layer
16.0 GT/s Capability is 0026h.
RO
31 24 23 16 15 8 7 0
PCI Express Extended Capability Header
16.0 GT/s Capabilities Register
Byte Offset
00h
04h
20h + Min (4, Maximum Link Width)
3Ch
16.0 GT/s Control Register 08h
16.0 GT/s Status Register
16.0 GT/s Local Data Parity Mismatch Status Register
0Ch
10h
16.0 GT/s First Retimer Data Parity Mismatch Status Register 14h
16.0 GT/s Second Retimer Data Parity Mismatch Status Register 18h
16.0 GT/s Reserved (RsvdP) 1Ch
20h
...
...
16.0 GT/s Lane Equalization Control Register
15 0
Next Capability Offset
31 20 19
Capability Version
PCI Express Extended Capability ID
16PCI Express Base Specification, Rev. 4.0 Version 1.0
826
Bit Location Register Description Attributes
19:16 Capability Version  This field is a PCI-SIG defined
version number that indicates the version of the Capability
structure present.
Must be 1h for this version of the specification.
RO
31:20 Next Capability Offset  This field contains the offset to the
next PCI Express Capability structure or 000h if no other
items exist in the linked list of Capabilities.
For Extended Capabilities implemented in Configuration
Space, this offset is relative to the beginning of
PCI-compatible Configuration Space and thus must always
be either 000h (for terminating list of Capabilities) or greater
than 0FFh.
RO
*/
func GetPCIeExtCapHeaderAddress(bus uint8, device uint8, function uint8, CapID uint16) uint16 {
	var address uint16 = 0x100
	// var address = ConfigReadu16(bus, device, function, uint16(startAddress))
	var value uint16
	var count int
	for true {
		value = ConfigReadu16(bus, device, function, address)
		if value == CapID {
			return address
		}
		address = ConfigReadu16(bus, device, function, address+0x2) >> 4
		if address == 0x0 {
			log.Fatal("Scanning to the end.")
			return ERRORu16
		}
		if count == 100 {
			log.Fatal("Cannot find match Ext CapID")
			break
		} else {
			count++
		}
	}
	return ERRORu16
}

/*
DisableHardwareAutonomousSpeed : Link Control 2 register
*/
func DisableHardwareAutonomousSpeed(bus uint8, device uint8, function uint8) {
	PCIeHeaderAddress := GetPCIeCapHeaderAddress(bus, device, function, 0x10)
	LinkControl2RegisterAddress := PCIeHeaderAddress + 0x30
	value := ConfigReadu16(bus, device, function, uint16(LinkControl2RegisterAddress))
	if value&(0x1<<5) == 0x0 {
		ConfigWriteu16(bus, device, function, uint16(LinkControl2RegisterAddress), uint16(value|(0x1<<5)))
		log.Printf("{0x%x:0x%x.0x%x} Disable hardware Autonomous Speed", bus, device, function)
	}
	// value = ConfigReadu16(bus, device, function, uint16(LinkControl2RegisterAddress))
	// log.Printf("AFTER\t - LinkControl2Register = 0x%x\n", value)
}

/*
DisableHardwareAutonomousWidth : Link Control register
*/
func DisableHardwareAutonomousWidth(bus uint8, device uint8, function uint8) {
	PCIeHeaderAddress := GetPCIeCapHeaderAddress(bus, device, function, 0x10)
	LinkControlRegisterAddress := PCIeHeaderAddress + 0x10
	value := ConfigReadu16(bus, device, function, uint16(LinkControlRegisterAddress))
	log.Printf("BEFORE\t - LinkControlRegister = 0x%x\n", value)
	if value&(0x1<<9) == 0x0 {
		ConfigWriteu16(bus, device, function, uint16(LinkControlRegisterAddress), uint16(value|(0x1<<9)))
		log.Printf("{0x%x:0x%x.0x%x} Disable hardware Autonomous Width", bus, device, function)
	}
	// value = ConfigReadu16(bus, device, function, uint16(LinkControlRegisterAddress))
	// log.Printf("AFTER\t - LinkControlRegister = 0x%x\n", value)
}

/*
GetHostBDF :
*/
func GetHostBDF(bus uint8, device uint8, function uint8) (uint8, uint8, uint8) {
	var SecondaryBusNumber uint8
	var SubordinateBusNumber uint8
	var EPBusNum = bus
	for RCBus := EPBusNum - 0x1; RCBus >= 0x0 && RCBus <= EPBusNum; RCBus-- {
		for RCDevice := uint8(0x0); RCDevice < 0x20; RCDevice++ {
			for RCFunction := uint8(0x0); RCFunction <= 0x8; RCFunction++ {
				SecondaryBusNumber = ConfigReadu8(RCBus,
					RCDevice, RCFunction, 0x19)
				SubordinateBusNumber = ConfigReadu8(RCBus,
					RCDevice, RCFunction, 0x1A)
				if SecondaryBusNumber == bus && SubordinateBusNumber == bus {
					return RCBus, RCDevice, RCFunction
				}
			}
		}

	}
	log.Fatal("Cannot Find the Root Complex for {%2x:%2x.%2x}", bus, device, function)
	return 0xFF, 0xFF, 0xFF
}

/*
CheckWhetherPCIeLinkSpeedReachGen4 :
*/
func CheckWhetherPCIeLinkSpeedReachGen4(bus uint8, device uint8, function uint8) bool {
	PCIeHeaderAddress := GetPCIeCapHeaderAddress(bus, device, function, 0x10)
	LinkStatusRegisterAddress := PCIeHeaderAddress + 0x12
	value := ConfigReadu8(bus, device, function, uint16(LinkStatusRegisterAddress))
	LinkSpeed := value & 0xF
	if LinkSpeed == 0x4 {
		return true
	} else {
		log.Printf("The Current Link Speed is Gen%x\n", LinkSpeed)
		return false
	}
}

/*
GetPcieLinkWidth :
*/
func GetPcieLinkWidth(bus uint8, device uint8, function uint8) uint8 {
	PCIeHeaderAddress := GetPCIeCapHeaderAddress(bus, device, function, 0x10)
	LinkStatusRegisterAddress := PCIeHeaderAddress + 0x12
	value := ConfigReadu8(bus, device, function, uint16(LinkStatusRegisterAddress))
	LinkWidth := value >> 4 & 0x3F
	return LinkWidth
}

/*
DisableASPM : Link Control register
*/
func DisableASPM(bus uint8, device uint8, function uint8) {
	PCIeHeaderAddress := GetPCIeCapHeaderAddress(bus, device, function, 0x10)
	LinkControlRegisterAddress := PCIeHeaderAddress + 0x10
	value := ConfigReadu8(bus, device, function, uint16(LinkControlRegisterAddress))

	if value&0x3 != 0x0 {
		ConfigWriteu8(bus, device, function, uint16(LinkControlRegisterAddress), uint8(value&0xFC))
		log.Printf("{0x%x:0x%x.0x%x} Disable ASPM", bus, device, function)
	}
}
