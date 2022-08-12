package main
import (
    "fmt"
    "bufio"
    "strings"
    "os"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        t:= sc.Text()
        a := strings.Split(t, ";")
        if a[0] == "S" {
            if a[1] == "M" {
                m := strings.Trim(a[2], "()")
                r := strings.SplitAfter(m, "")
                for _, v := range r {
                    
                }
            }
        } else if a[0] == "C" {
            
        }
    }
}