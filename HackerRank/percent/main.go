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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var countP = 0
	var countN = 0
	var countZ = 0

	var total = len(arr)

	for _, v := range arr {
		if int64(v) > 0 {
			countP++
			continue
		}

		if int64(v) < 0 {
			countN++
			continue
		}

		if int64(v) == 0 {
			countZ++
		}
	}

	totalS := float64(total)
	percentP := float64(countP) / totalS
	percentN := float64(countN) / totalS
	percentZ := float64(1 - percentP - percentN)

	fmt.Println(roundFloat(percentP, 6))
	fmt.Println(roundFloat(percentN, 6))
	fmt.Println(roundFloat(percentZ, 6))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
