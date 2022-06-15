package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'icecreamParlor' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER m
 *  2. INTEGER_ARRAY arr
 */

/*
 F(N) =    max[f(j)] + T
      =  max[fj)
( 0 <= j <= N)


 have match t: t = 1
 dont have match t : t = 0


*/

func longestFind(m int, arr []int) {
	// Write your code here
	fmt.Println(m, arr)
	// array array save (t,m) in index money in this select
	result := make([]int, m+1)
	trace := make([]int, m+1)

	// init array
	for i := 0; i < m+1; i++ {
		result[i] = -1
		trace[i] = -1
	}

	// start point DP: m = 0
	for i := 0; i < m+1; i++ {
		result[0] = 0
		trace[0] = 0
	}

	fmt.Println("result =", result, "trace=", trace)

	//DP
	for i := 1; i < m+1; i++ {
		//if have select
		temp := i
		for trace[temp] == -1 {
			temp--
		}

		if trace[temp] < arr[i-1] {
			result[i] = result[i-1] + 1
			trace[i] = arr[i-1]
		} else {
			result[i] = result[i-1]
			//trace[i] = -1
		}

		////// if i , check i-1 to 1;   tim max
		//check := arr[i-1]
		//listResult := make([]int, m)
		//listTrace := make([]int, m)
		//
		//for t := i; t > 0; t-- {
		//	if trace[t-1] != -1 {
		//		if check > trace[t-1] {
		//			resultT := result[t-1] + 1
		//			listResult[t-1] = resultT
		//			listTrace[t-1] = trace[t-1]
		//		}
		//	}
		//}
		//
		//fmt.Println("listResult =", listResult, "listTrace=", listTrace)
		//
		//maxTemp := listResult[0]
		//traceMin := trace[0]
		//for i, j := range listResult {
		//	if j > maxTemp {
		//		maxTemp = j
		//	} else if j == maxTemp {
		//		if listTrace[i] < traceMin {
		//			traceMin = listTrace[i]
		//		}
		//	}
		//}
		//
		//result[i] = maxTemp
		//trace[i] = check

	}

	fmt.Println("result =", result, "trace=", trace)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var arr []int

	for i := 0; i < int(t); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}
	longestFind(t, arr)
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
