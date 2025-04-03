// functions, multiple returns, conditionals, loops
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(add(42, 13))
	fmt.Println("Men's team: " + basketballmessage(65, 68))   // men's
	fmt.Println("Women's team: " + basketballmessage(56, 68)) // women's

	msg, err := maybefaultybasketballmessage(65, -68)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	msg, err = maybefaultybasketballmessage(56, 68)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	visualizeourscore(56)

	fmt.Printf("%s\n", switchbbmessage(56, 68))
}

func add(x int, y int) int {
	return x + y
}

func basketballmessage(ourscore uint, theirscore uint) string {
	if ourscore > theirscore {
		return "We won!"
	} else if theirscore-ourscore < 5 {
		return "Wow we almost won!"
	} else {
		return "We didn't win but played great this year!"
	}
}

func maybefaultybasketballmessage(ourscore int, theirscore int) (string, error) {
	var err error

	if ourscore < 0 || theirscore < 0 {
		err = errors.New("The score cannot be negative")
		return "", err
	} else {
		return basketballmessage(uint(ourscore), uint(theirscore)), nil
	}
}

func visualizeourscore(ourscore int) {
	// https://tinyurl.com/tritonsymbol

	for i := 0; i < ourscore; i++ {
		fmt.Printf("🔱")
	}
	fmt.Printf("\n")
}

func switchbbmessage(ourscore uint, theirscore uint) string {
	switch {
	case ourscore > theirscore:
		return "We won!"
	case theirscore-ourscore < 5:
		return "Wow we almost won!"
	default:
		return "We didn't win but played great this year!"
	}
}
