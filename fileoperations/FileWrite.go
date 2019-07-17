// FileWrite.go File

// This file was mostly used for learning purpose.
// Most of the code is from https://colobu.com/2016/10/12/go-file-operations/
// I don't know how many api will be used in further coding.
// -- Chi Zhang 7/6/2019

//Package fileoperations provides some sample api for file operations.
package fileoperations

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

// WriteFile is a normal way of writing file
func WriteFile(fileName string, content string) bool {
	file, err := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// translate string to bytes and put them into an array
	byteSlice := []byte(content)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
	return true
}

// FastWriteFile uses iotuil to operate(create/open) files
func FastWriteFile(fileName string, content string) bool {
	err := ioutil.WriteFile(fileName,
		[]byte(content),
		0666)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// BufWriteFile uses bufio package and all the bytes were written to buffer before written into hard disk.
// It will speed up the IO performance.
func BufWriteFile(fileName string, content string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer file.Close()
	// Create a buffer for the file
	bufferedWriter := bufio.NewWriter(file)
	// Write bytes into buffer
	// bytesWritten, err := bufferedWriter.Write(
	// 	[]byte{65,66,67},
	// )
	// There are also WriteRune() and WriteByte() functions

	// Write String into buffer
	bytesWritten, err := bufferedWriter.WriteString(
		content,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// Check the bytes in buffer
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// Check the available bytes in buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Available buffered: %d\n", bytesAvailable)

	// Flush the bytes from buffer into hard disk
	bufferedWriter.Flush()

	// Reset buffer
	bufferedWriter.Reset(bufferedWriter)
	
	bytesAvailable = bufferedWriter.Available()
	log.Printf("Available buffered: %d\n", bytesAvailable)

	// Set the size of buffer (increase the size)
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)
	
	bytesAvailable = bufferedWriter.Available()
	log.Printf("Available buffered: %d\n", bytesAvailable)
	return true
}
