/*
Given an array of integers and a positive integer k,
determine the number of (i,j) pairs where i < j and ar[i] + ar[j] is divisible by k.

Example

ar = [1,2,3,4,5,6]
k = 5


Three pairs meet the criteria:  and .

[1,4], [2,3], and [4,6].

Function Description

Complete the divisibleSumPairs function in the editor below.

divisibleSumPairs has the following parameter(s):

int n: the length of array
int ar[n]: an array of integers
int k: the integer divisor

Returns
- int: the number of pairs

Input Format

The first line contains 2 space-separated integers,
n  and k.
The second line contains n space-separated integers,
 each a value of arr[i].

Constraints

2 < n < 100
1 < k < 100
1 < ar[i] < 100

Sample Input

STDIN           Function
-----           --------
6 3             n = 6, k = 3
1 3 2 6 1 2     ar = [1, 3, 2, 6, 1, 2]
Sample Output

 5
Explanation

Here are the 5 valid pairs when

k = 3

(0,2) -> ar[0] + ar[2] = 1 + 2 = 3
(0, 5) -> ar[0] + ar[5] = 1 + 2 = 3
(1, 3) -> ar[1] + ar[3] = 3 + 6 = 9
(2, 4) -> ar[2] + ar[4] = 2 + 1 = 3
(4, 5) -> ar[4] + ar[5] = 1 + 2 = 3

*/

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
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var (
		i, j, c int32
	)
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				c++
			}
		}
	}
	return c
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := divisibleSumPairs(n, k, ar)

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
