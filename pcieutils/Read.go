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
