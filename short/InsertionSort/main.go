package main

import "fmt"

func insert(a []int, index int, value int) []int {
	a = append(a[:index+1], a[index:]...) // Step 1+2
	a[index] = value                      // Step 3
	return a
}
func InsertionShort(input []int) []int {
	var output []int
	for i := 0; i < len(input); i++ {
		if len(output) == 0 {
			output = append(output, input[i])
		} else {
			find := len(output) - 1
			for j := 0; j < len(output); j++ {
				//insert
				if input[i] < output[j] {
					find = j
					output = insert(output, find, input[i])
					break
				}

				//addpend to end
				if j == len(output)-1 {
					output = append(output, input[i])
					break
				}
			}
		}

	}
	return output
}

func main() {
	a := []int{20, 3, 213, 2, 1, 3, 5, 342, 2, 1}
	fmt.Println(a)
	fmt.Println(InsertionShort(a))

}

/**
Big -O O(n2)
*/
