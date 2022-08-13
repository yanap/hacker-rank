package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ToSpaceDelimited(s string) (out string) {
	for _, v := range s {
		if unicode.IsUpper(v) {
			l := unicode.ToLower(v)
			out += " " + string(l)
		} else {
			out += string(v)
		}
	}
	return out
}

func ToCamel(s, ty string) (out string) {
	de := strings.Split(s, " ")
	for i, w := range de {
		for j, v := range w {
			if ty == "C" && i == 0 && j == 0 {
				u := unicode.ToUpper(v)
				out += string(u)
			} else if i > 0 && j == 0 {
				u := unicode.ToUpper(v)
				out += string(u)
			} else {
				out += string(v)
			}
		}
	}
	return out
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	sc := bufio.NewScanner(os.Stdin)
	var out string
	for sc.Scan() {
		t := sc.Text()
		a := strings.Split(t, ";")
		if a[0] == "S" {
			if a[1] == "M" {
				m := strings.Trim(a[2], "()")
				out = ToSpaceDelimited(m)
			} else if a[1] == "C" {
				out = ToSpaceDelimited(a[2])
			} else if a[1] == "V" {
				out = ToSpaceDelimited(a[2])
			}
		} else if a[0] == "C" {
			if a[1] == "M" {
				out = ToCamel(a[2], a[1])
				out += "()"
			} else if a[1] == "C" {
				out = ToCamel(a[2], a[1])
			} else if a[1] == "V" {
				out = ToCamel(a[2], a[1])
			}
		}
		fmt.Println(out)
	}
}
