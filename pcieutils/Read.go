// Read.go File

// This file was mostly used for Config/Mem space read access
// -- Chi Zhang 7/10/2019

//Package pcieutils provides some api for Config/Mem space read access.
package pcieutils

import (

	// "bufio"
	"log"
	"os"
	"syscall"
	"unsafe"
)

//BaseAddress : local parameter
var BaseAddress int64

func init() {
	BaseAddress = int64(GetPCIeBaseAddress())
	// fmt.Printf("BaseAddress = 0x%x\n", BaseAddress)
}

// Readu32 is for read Config space and return 32 bits value
func Readu32(baseAddress int64, offset uint16) uint32 {
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
	value = *(*uint32)(unsafe.Pointer(&mmap[offset]))
	err = syscall.Munmap(mmap)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// ConfigReadu32 is for read device's Config Space and return 32 bits value
func ConfigReadu32(bus uint8, device uint8, function uint8, offset uint16) uint32 {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	value := Readu32(address, offset)
	return value
}

// Readu16 is for read Config space and return 16 bits value
func Readu16(baseAddress int64, offset uint16) uint16 {
	var value uint16 = 0xFFFF
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
	value = *(*uint16)(unsafe.Pointer(&mmap[offset]))
	err = syscall.Munmap(mmap)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// ConfigReadu16 is for read device's Config Space and return 16 bits value
func ConfigReadu16(bus uint8, device uint8, function uint8, offset uint16) uint16 {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	value := Readu16(address, offset)
	return value
}

// Readu8 is for read Config space and return 8 bits value
func Readu8(baseAddress int64, offset uint16) uint8 {
	var value uint8 = 0xFF
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
	value = *(*uint8)(unsafe.Pointer(&mmap[offset]))
	err = syscall.Munmap(mmap)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// ConfigReadu8 is for read device's Config Space and return 8 bits value
func ConfigReadu8(bus uint8, device uint8, function uint8, offset uint16) uint8 {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	value := Readu8(address, offset)
	return value
}
