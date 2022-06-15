package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(5)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quickSort(slice, 0, len(slice)-1)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

//func quicksort1(a []int, left int, right int) {
//	if left >= right {
//		return
//	}
//
//	r := right - left + 1
//	pivot := rand.Intn(r)
//	vPivot := a[pivot]
//	i := left
//	j := right
//
//	for {
//		// find first element to  >= pivot
//		for a[i] < vPivot {
//			i++
//		}
//
//		//find end element to <= pivot
//		for a[j] > vPivot {
//			j--
//		}
//
//		// switch
//		if i <= j {
//			if i < j {
//				temp := a[i]
//				a[i] = a[j]
//				a[j] = temp
//			}
//			i++
//			j--
//		}
//
//		if i > j {
//			break
//		}
//	}
//
//	quicksort(a, left, j)
//	quicksort(a, i, right)
//}
