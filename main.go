package main

import (
	// "time"
	// "PLS_GO/fileoperations"
	"PLS_GO/pcieutils"
	"fmt"
)

func main() {
	// fmt.Println("hello world.")
	// FileName := "test_create_file1.log"
	// configfile.CreateNewEmptyFile(FileName)
	// configfile.TruncateFile(FileName, 50)
	// configfile.GetFileInfo(FileName)
	// fileoperations.BufWriteFile(FileName, "Hello World!!")
	// baseAdr := int64(0xF8000000 + 0x1e*0x100000)

	// value1 := pcieutils.Readu32(baseAdr, 0xf0)
	// fmt.Println("mask0")
	// value1 = value1&0xFFFFFFF0 | 0x3
	// fmt.Println("mask01")
	// pcieutils.Writeu16(baseAdr, 0xf0, uint16(value1))

	value1 := pcieutils.GetPCIeBaseAddress()
	// value2 := pcieutils.Readu32(baseAdr, 0x104)
	// value3 := pcieutils.Readu32(int64(0xF7900000), 0x28)
	// value4 := pcieutils.Readu32(int64(0xF7900000), 0x2C)
	// value5 := pcieutils.Readu32(int64(0xffffe000), 0x0)
	fmt.Printf("Value1 = 0x%x\n", value1)
	// fmt.Printf("Value2 = 0x%x\n", value2)
	// fmt.Printf("Value3 = 0x%x\n", value3)
	// fmt.Printf("Value4 = 0x%x\n", value4)
	// fmt.Printf("Value5 = 0x%x\n", value5)
	// go pcieutils.Polling(baseAdr, 0x110, 30, 5)
	// go pcieutils.Polling(baseAdr, 0x104, 50, 5)
	// time.Sleep(6 * time.Second)
}
