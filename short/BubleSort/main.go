package main

import "fmt"

func BubleShort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := len(input) - 1; j > i; j-- {
			if input[j] < input[j-1] {
				temp := input[j]
				input[j] = input[j-1]
				input[j-1] = temp
			}
		}
	}
}

func main() {
	a := []int{20, 3, 213, 2, 1, 3, 5, 342, 2, 1}
	fmt.Println(a)
	BubleShort(a)
	fmt.Println(a)
}

/**
O(N2)
*/
