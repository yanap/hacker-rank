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
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var l, r, s int32
	for i, n := range arr {
		fmt.Println("i", i)
		for j, m := range n {
			fmt.Println("j", j)
			if i == 0 && j == 0 {
				l = m
				fmt.Println("00m", m)
				fmt.Println("00l", l)
			}
			if i == 0 && j == 2 {
				r = m
				fmt.Println("02m", m)
				fmt.Println("02r", r)
			}
			if i == 1 && j == 1 {
				l = l + m
				r = r + m
				fmt.Println("11m", m)
				fmt.Println("11l", l)
				fmt.Println("11r", r)
			}
			if i == 2 && j == 0 {
				r = r + m
				fmt.Println("20m", m)
				fmt.Println("20r", r)
			}
			if i == 2 && j == 2 {
				l = l + m
				fmt.Println("22m", m)
				fmt.Println("22l", l)
			}
		}
	}
	fmt.Println("l", l)
	fmt.Println("r", r)
	if l > r {
		s = l - r
	} else {
		s = r - l
	}
	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
