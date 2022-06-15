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
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	l := len(s)
	var sum int32
	sum = 0
	for i := 0; i <= l-1; i++ {
		for jumpLen := 0; jumpLen <= l-1; jumpLen++ {
			countAnagrams(s, i, jumpLen, &sum)
		}
	}

	return sum
}
func countAnagrams(s string, headStart int, lenJump int, count *int32) {
	l := len(s)
	if headStart < 0 || headStart > l-1 {
		return
	}

	for i := headStart; i <= (l-1)-lenJump+1; i++ {
		var firstString string
		var secondString string
		for j := i + 1; j <= (l-1)-lenJump-1; j++ {
			if lenJump != 0 {
				firstString = string(s[i : i+lenJump-1])
				secondString = string(s[j : j+lenJump-1])
			} else {
				firstString = string(s[i])
				secondString = string(s[j])
			}

			if isAnagrams(firstString, secondString) {
				*count++
			}
		}
	}
}

func isAnagrams(s string, t string) bool {

	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}

//func isAnagrams(firstString string, secondString string) bool {
//	if len(firstString) == 0 || len(secondString) == 0 {
//		return false
//	}
//
//	if len(firstString) != len(secondString) {
//		return false
//	}
//
//	for i := 0; i < len(firstString)-1; i++ {
//		if !strings.Contains(string(firstString[i]), secondString) {
//			return false
//		}
//	}
//
//	for i := 0; i < len(secondString)-1; i++ {
//		if !strings.Contains(string(secondString[i]), firstString) {
//			return false
//		}
//	}
//
//	return true
//}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	os.Setenv("OUTPUT_PATH", "out.txt")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)
		fmt.Println(result)
		fmt.Fprintf(writer, "%d\n", result)
	}

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
