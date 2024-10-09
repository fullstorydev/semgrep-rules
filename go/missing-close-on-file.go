package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "example.txt"

	err := readFileAndPrint(filename)
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = readFileAndPrint_fp(filename)
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createTempFile()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createTempFile_fp()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createTempFile2()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createTempFile2_fp()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createSimpleFile()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}

	err = createSimpleFile_fp()
	if err != nil {
		log.Fatal("Error reading and printing file:", err)
	}
}

func readFileAndPrint(filename string) error {
	// ruleid: missing-close-on-file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(content))

	return nil
}

func readFileAndPrint_fp(filename string) error {
	// ruleid: missing-close-on-file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(content))

	return nil
}

func createTempFile() error {
	// ruleid: missing-close-on-file
	tempFile, err := ioutil.TempFile("", "example-*.txt")
	if err != nil {
		return err
	}

	content := []byte("test")
	_, err = tempFile.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Temporary file created:", tempFile.Name())

	return nil
}

func createTempFile_fp() error {
	// ok: missing-close-on-file
	tempFile, err := ioutil.TempFile("", "example-*.txt")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	content := []byte("test")
	_, err = tempFile.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Temporary file created:", tempFile.Name())

	return nil
}

func createTempFile2() error {
	// ruleid: missing-close-on-file
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		return err
	}

	content := []byte("test")
	_, err = tempFile.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Temporary file created:", tempFile.Name())

	return nil
}

func createTempFile2_fp() error {
	// ok: missing-close-on-file
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	content := []byte("test")
	_, err = tempFile.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Temporary file created:", tempFile.Name())

	return nil
}

func createSimpleFile() error {
	fileName := "example.txt"

	// ruleid: missing-close-on-file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	content := []byte("test")
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("file created:", fileName)

	return nil
}

func createSimpleFile_fp() error {
	fileName := "example.txt"

	// ok: missing-close-on-file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	file.Close()

	content := []byte("test")
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("file created:", fileName)

	return nil
}
