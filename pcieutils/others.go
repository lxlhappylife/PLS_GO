// others.go File

// This file was mostly used for pcie access related features
// -- Chi Zhang 7/10/2019

//Package pcieutils provides some api for Config/Mem space read access.
package pcieutils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

var counterForOneMicrosecond int64

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

	fmt.Printf("%s", buf)
	return address
}
