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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	slc := strings.Split(s, ":")
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}

	timeType := slc[2][2:]
	//fmt.Println("timeType", timeType)
	if timeType == "PM" {
		h, _ := strconv.Atoi(slc[0])
		h += 12
		if h == 24 {
			slc[0] = "12"
		} else {
			slc[0] = strconv.Itoa(h)
		}
	}

	if timeType == "AM" {
		h, _ := strconv.Atoi(slc[0])
		if h == 12 {
			slc[0] = "00"
		}
	}

	fmt.Println(slc[0] + ":" + slc[1] + ":" + slc[2][0:2])
	return slc[0] + ":" + slc[1] + ":" + slc[2][0:2]
}

func main() {
	//fmt.Println("start main")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	os.Setenv("OUTPUT_PATH", "time.txt")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
