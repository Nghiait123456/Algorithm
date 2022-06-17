package main
import (
	"fmt"
	"io"
	"bufio"
	"strconv"
	"strings"
	"os"
	"math"
)

var traceTable []int
func dpJump(a []int)  {
	fmt.Println(a)

	//init traceTable for start point
	fmt.Println(traceTable)
	traceTable = append(traceTable, 0)
	traceTable = append(traceTable, AbsInt(a[0] - a[1]))
	fmt.Println(traceTable)

	//rule DP
	/**
	 fi  =  min( (f(i-1) + abs( i-1, i)) and ( f(u-2) + abs(i-2,i)))
	 */

	for i := 2; i < len(a) ; i ++ {
		firstWay := traceTable[i-1] + AbsInt(a[i] - a[i-1])
		secondWay := traceTable[i-2] + AbsInt(a[i] - a[i-2])

		fmt.Println("i= ", i, "ai= ", a[i], "firstWay= ", firstWay, "secondWay", secondWay)

		if firstWay < secondWay {
			traceTable = append(traceTable, firstWay)
		} else {
			traceTable = append(traceTable, secondWay)
		}
	}

	fmt.Println(traceTable)
}

func AbsInt(n int) int {
	return int(math.Abs(float64(n)))
}
func main()  {
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

	dpJump(a)
	//for _, v := range globalResult {
	//	fmt.Printf("%.1f\n", v)
	//}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
