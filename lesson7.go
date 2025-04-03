package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	filename := "demo.bin"
	data := []byte{0x01, 0x02, 0xA3, 0xFF, 0x00}

	// Write bytes to file
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read bytes back from file
	readData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Compare for equality
	if bytes.Equal(data, readData) {
		fmt.Println("Success! Data matches.")
	} else {
		fmt.Println("Data mismatch!")
	}
}
