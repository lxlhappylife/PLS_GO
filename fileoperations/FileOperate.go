// FileOperate.go File

// This file was mostly used for learning purpose.
// I don't know how many api will be used in further coding.
// -- Chi Zhang 7/6/2019

//Package fileoperations provides some sample api for file operations.
package fileoperations

import (
	"fmt"
	"log"
	"os"
	"io"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("current path is %s.\n", path)
}

// CreateNewEmptyFile is for Create Files
func CreateNewEmptyFile(Filename string) bool {
	// fmt.Printf("Filename=%s", Filename)
	newFile, err := os.Create(Filename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println(newFile)
	newFile.Close()
	return true
}

// TruncateFile is for truncate Files
func TruncateFile(Filename string, size int64) bool {
	err := os.Truncate(Filename, size)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetFileInfo is for Return File Infomation
func GetFileInfo(Filename string) {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(Filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File name:", fileInfo.Name())
	log.Println("Size in bytes:", fileInfo.Size())
	log.Println("Permissions:", fileInfo.Mode())
	log.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

// CopyFiles is for Copy old file to new file
func CopyFiles(oldFilename string, newFilename string) bool {
	originalFile, err := os.Open(oldFilename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer originalFile.Close()

	// Create New Files as target file
	newFile, err := os.Create(newFilename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer newFile.Close()

	// Copy the bytes from resource to target file
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Printf("Copied %d bytes.", bytesWritten) 

	// flush the content into hard disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
