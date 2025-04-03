// file IO and 'defer'
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type mapposition struct {
	name      string
	lat, long float64
	height    int32
}

func writemapdata() {
	// Create or open the file (overwrites existing file or creates a new one)
	file, err := os.Create("mapdata.dat")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed after operation

	// Create a buffered writer for efficiency
	writer := bufio.NewWriter(file)

	mappositions := []mapposition{
		{"UC San Diego", 32.8812, 117.2344, 25},
		{"UC Berkeley", 37.8712, 122.2555, 51},
	}

	// Write the positions to a file
	for _, pos := range mappositions {
		writer.WriteString(pos.name + "\n")
		err := binary.Write(writer, binary.LittleEndian, pos.lat)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
		err = binary.Write(writer, binary.LittleEndian, pos.long)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
		err = binary.Write(writer, binary.LittleEndian, pos.height)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}

	// Ensure all buffered operations are applied
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

	fmt.Println("File written successfully.")
}

func readmapdata() {
	file, err := os.Open("mapdata.dat")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for i := 0; i < 2; i++ {
		var pos mapposition

		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		pos.name = strings.TrimSpace(str)

		err = binary.Read(reader, binary.LittleEndian, &pos.lat)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}

		err = binary.Read(reader, binary.LittleEndian, &pos.long)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}

		err = binary.Read(reader, binary.LittleEndian, &pos.height)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}

		fmt.Printf("name: %v, lat: %v, long: %v, height: %v\n", pos.name, pos.lat, pos.long, pos.height)
	}
}

func main() {
	writemapdata()
	readmapdata()
}
