package main

import "fmt"

func filter(nums []int, f func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	evens := filter(numbers, isEven)
	fmt.Println("Even numbers:", evens)
	odds := filter(numbers, isOdd)
	fmt.Println("Odd numbers:", odds)
}
