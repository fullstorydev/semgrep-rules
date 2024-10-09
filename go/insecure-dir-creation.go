package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	createFile()
	createFile_fp()
	createFile2()
	createFile2_fp()
	createDir()
	createDir_fp()
}

func createFile() {
	d1 := []byte("hello\ngo again\n")

	_, err := os.Stat("./dat1")
	if err != nil {
		fmt.Println(err)
	}
	// ruleid: insecure-dir-creation
	err = os.WriteFile("./dat1", d1, 0700)
	if err != nil {
		fmt.Println(err)
	}
}

func createFile2() {
	d1 := []byte("hello\ngo again\n")

	_, err := os.Stat("./dat1")
	if err != nil {
		fmt.Println(err)
	}
	// ruleid: insecure-dir-creation
	err = ioutil.WriteFile("./dat1", d1, 0700)
	if err != nil {
		fmt.Println(err)
	}
}

func createDir() {
	dirName := "temp"

	// ruleid: insecure-dir-creation
	err := os.Mkdir(dirName, 0700)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dirName, err)
		return
	}

	fmt.Printf("Directory %s created successfully.\n", dirName)
}

func createFile_fp() {
	d1 := []byte("hello\ngo again\n")

	if _, err := os.Stat("./dat1"); os.IsNotExist(err) {
		// ok: insecure-dir-creation
		err = os.WriteFile("./dat1", d1, 0700)
		if err != nil {
			fmt.Println(err)
		}
	}
}
