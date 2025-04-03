// variables, types
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)

	var i int = 3
	fmt.Println(i)
	i = i + 2
	fmt.Println(i)

	var smallnum uint8
	fmt.Println(smallnum)
	smallnum = 255
	fmt.Println(smallnum)
	smallnum += 1
	fmt.Println(smallnum)

	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte (uint8)
	// rune
	// float32 float64

	var greeting string = "Welcome!"
	fmt.Println(greeting)

	var emptygreeting string
	fmt.Println(emptygreeting)

	shortgreeting := "yo!"
	fmt.Println(shortgreeting)

	var mybool bool = true
	fmt.Println(mybool)
	mybool = false
	fmt.Println(mybool)

	greeting, mybool = "new greeting", true
	fmt.Printf("%s %t\n", greeting, mybool)
}
