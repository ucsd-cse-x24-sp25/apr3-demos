// structs, arrays, slices, and maps
package main

import (
	"fmt"
	"slices"
)

type GameOutcome struct {
	teamname   string
	ourscore   uint
	theirscore uint
}

func main() {
	mens := GameOutcome{
		teamname:   "Men's basketball",
		ourscore:   65,
		theirscore: 68,
	}

	womens := GameOutcome{
		teamname:   "Women's basketball",
		ourscore:   56,
		theirscore: 68,
	}

	printscore(mens)
	printscore(womens)

	var simplearray [5]int32
	simplearray[0] = 3
	simplearray[1] = 5
	simplearray[2] = 4
	simplearray[3] = 2
	simplearray[4] = 1

	fmt.Println(simplearray[4])
	fmt.Println(simplearray[2:4])
	fmt.Println(simplearray[:])

	slices.Sort(simplearray[:])
	fmt.Println(simplearray[:])

	var numslice []int32 = make([]int32, 5)
	numslice[3] = 2

	for idx, val := range numslice {
		fmt.Println(idx, val)
	}

	for _, val := range numslice {
		fmt.Println(val)
	}

	scores := []GameOutcome{mens, womens}
	for _, score := range scores {
		printscore(score)
	}

	schoolages := map[string]uint{
		"ucsd": 64,
		"ucb":  157,
		"ucla": 98,
	}
	fmt.Println(schoolages["ucb"])

	schools := make([]string, 0)
	for name, _ := range schoolages {
		schools = append(schools, name)
	}
	fmt.Println(schools)

	schoolages2 := make(map[string]uint)
	schoolages2["merced"] = 19
	fmt.Println(schoolages2["merced"])
}

func printscore(outcome GameOutcome) {
	fmt.Printf("%v: (%d, %d)\n", outcome.teamname, outcome.ourscore, outcome.theirscore)
}
