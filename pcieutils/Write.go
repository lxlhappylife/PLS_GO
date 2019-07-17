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
