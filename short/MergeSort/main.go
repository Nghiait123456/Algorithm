package main

import "fmt"

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	left := MergeSort(input[:(len(input) / 2)])
	right := MergeSort(input[(len(input) / 2):])

	return Merger(left, right)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func Merger(left []int, right []int) []int {
	var out []int
	indexLeft := 0
	indexRight := 0

	for {
		if len(left) == 0 && len(right) == 0 {
			break
		}

		if len(left) == 0 && len(right) != 0 {
			out = append(out, right[indexRight:]...)
			break
		}

		if len(right) == 0 && len(left) != 0 {
			out = append(out, left[indexLeft:]...)
			break
		}

		if left[indexLeft] > right[indexRight] {
			out = append(out, right[indexRight])
			right = remove(right, 0)
		} else if left[indexLeft] < right[indexRight] {
			out = append(out, left[indexLeft])
			left = remove(left, 0)
		} else if left[indexLeft] == right[indexRight] {
			out = append(out, left[indexLeft])
			out = append(out, right[indexRight])
			right = remove(right, 0)
			left = remove(left, 0)
		}
	}

	//fmt.Printf("start merger, left = %v, right = %v, out = %v", left, right, out)
	return out
}

func main() {
	a := []int{20, 3, 213, 2, 1, 3, 5, 342, 2, 1}
	fmt.Println(a)
	fmt.Println(MergeSort(a))
}

/**
O(N∗logN)
*/
