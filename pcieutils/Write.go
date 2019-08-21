// Write.go File

// This file was mostly used for Config/Mem space read access
// -- Chi Zhang 7/10/2019

//Package pcieutils provides some api for Config/Mem space read access.
package pcieutils

import (
	"log"
	"os"
	"syscall"
	"unsafe"
)

// Writeu32 is for write Config space
func Writeu32(baseAddress int64, offset uint16, value uint32) {
	const bufferSize int = 4096
	file, err := os.OpenFile("/dev/mem", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mmap, err := syscall.Mmap(int(file.Fd()), baseAddress, bufferSize, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	*(*uint32)(unsafe.Pointer(&mmap[offset])) = value
	err = syscall.Munmap(mmap)

	if err != nil {
		log.Fatal(err)
	}
}

// ConfigWriteu32 is for Write device's Config Space with 32 bits value
func ConfigWriteu32(bus uint8, device uint8, function uint8, offset uint16, value uint32) {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	Writeu32(address, offset, value)

}

// Writeu16 is for write Config space
func Writeu16(baseAddress int64, offset uint16, value uint16) {
	const bufferSize int = 4096
	file, err := os.OpenFile("/dev/mem", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mmap, err := syscall.Mmap(int(file.Fd()), baseAddress, bufferSize, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	*(*uint16)(unsafe.Pointer(&mmap[offset])) = value
	err = syscall.Munmap(mmap)

	if err != nil {
		log.Fatal(err)
	}
}

// ConfigWriteu16 is for Write device's Config Space with 16 bits value
func ConfigWriteu16(bus uint8, device uint8, function uint8, offset uint16, value uint16) {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	Writeu16(address, offset, value)

}

// Writeu8 is for write Config space
func Writeu8(baseAddress int64, offset uint16, value uint8) {
	const bufferSize int = 4096
	file, err := os.OpenFile("/dev/mem", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mmap, err := syscall.Mmap(int(file.Fd()), baseAddress, bufferSize, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	*(*uint8)(unsafe.Pointer(&mmap[offset])) = value
	err = syscall.Munmap(mmap)

	if err != nil {
		log.Fatal(err)
	}
}

// ConfigWriteu8 is for Write device's Config Space with 16 bits value
func ConfigWriteu8(bus uint8, device uint8, function uint8, offset uint16, value uint8) {
	address := BaseAddress | int64(bus)<<(4*5) | int64(device)<<15 | int64(function)<<12
	// fmt.Printf("address = 0x%X\n", address)
	Writeu8(address, offset, value)

}
