package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'runningMedian' function below.
 *
 * The function is expected to return a DOUBLE_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

var globalResult []float64
var globalInput []int

//func partition(arr []int, low, high int) ([]int, int) {
//	pivot := arr[high]
//	i := low
//	for j := low; j < high; j++ {
//		if arr[j] < pivot {
//			arr[i], arr[j] = arr[j], arr[i]
//			i++
//		}
//	}
//	arr[i], arr[high] = arr[high], arr[i]
//	return arr, i
//}
//
//func quickSort(arr []int, low, high int) []int {
//	if low < high {
//		var p int
//		arr, p = partition(arr, low, high)
//		arr = quickSort(arr, low, p-1)
//		arr = quickSort(arr, p+1, high)
//	}
//	return arr
//}

func runningMedian(a []int) {
	// Write your code here
	for _, j := range a {
		lenGolebalInput := len(globalInput)

		//insert and short
		if lenGolebalInput == 0 {
			globalInput = append(globalInput, j)
			//fmt.Println("globalInput len =0", globalInput)
		} else if lenGolebalInput > 0 {
			find := 0
			for i, v := range globalInput {
				if i < lenGolebalInput-1 {
					if j <= v {
						//fmt.Println("before change", "i", i, "j", j, "globalInput", globalInput, "globalInput[:i]", globalInput[:i], "globalInput[i:]", globalInput[i:])
						globalInput = append(globalInput[:i+1], globalInput[i:]...)
						globalInput[i] = j
						//fmt.Println("after change", "i", i, "j", j, "globalInput", globalInput, "globalInput[:i]", globalInput[:i], "globalInput[i:]", globalInput[i:])
						find = 1
						break
					}
				} else if i == lenGolebalInput-1 {
					if j <= v {
						globalInput = append(globalInput, globalInput[i])
						globalInput[i] = j
						find = 1
						break
					}
				}
			}

			if find == 0 {
				globalInput = append(globalInput, j)
			}
		}

		//fmt.Println(globalInput)
		//get medium
		lenAShot := len(globalInput)
		if lenAShot%2 == 0 {
			if lenAShot > 2 {
				mid := lenAShot / 2
				result := math.Round((float64(globalInput[mid-1]+globalInput[mid])/2)*10) / 10
				globalResult = append(globalResult, result)
				//fmt.Printf("%.1f\n", result)
			} else if lenAShot == 2 {
				result := math.Round((float64(globalInput[0]+globalInput[1])/2)*10) / 10
				//fmt.Printf("%.1f\n", result)
				globalResult = append(globalResult, result)
			}
		} else if lenAShot%2 != 0 {
			index := (lenAShot+1)/2 - 1
			result := float64(globalInput[index])
			globalResult = append(globalResult, result)
			//fmt.Printf("%.1f\n", result)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	var a []int

	for i := 0; i < int(aCount); i++ {
		aItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		a = append(a, aItem)
	}

	runningMedian(a)
	for _, v := range globalResult {
		fmt.Printf("%.1f\n", v)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
