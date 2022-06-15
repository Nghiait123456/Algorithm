package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fibonaci(max int32) {
	result := make([]int32, int(max))
	result[1] = 1
	result[2] = 1

	for i := 3; int32(i) < max; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	fmt.Println(result)
}

func checkArrayHaveValue(value int32, arr []int32) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	fibonaci(t)

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
