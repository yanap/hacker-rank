package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		t := sc.Text()
		a := strings.Split(t, ";")
		if a[0] == "S" {
			if a[1] == "M" {
				m := strings.Trim(a[2], "()")
				for _, v := range m {
					if unicode.IsUpper(v) {
					} else {
					}
				}
			}
		} else if a[0] == "C" {

		}
	}
}
