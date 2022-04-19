package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "unicode"
)

func split(s string) string {
    outStr := ""
    for _, c := range s {
        if outStr == "" { // first character
            outStr += string(unicode.ToLower(c))
        } else {
            if unicode.IsUpper(c) { // find a hump, add a space delimit
                outStr += " "
                outStr += string(unicode.ToLower(c))
            } else {
                outStr += string(c)
            }
        }
    }
    return outStr
}

func combine(t string, s string) string {
    outStr := ""
    hump := false
    for _, c := range s {
        if outStr == "" {  // first character
            if t == "C" { // Class need to capitalize
                outStr += string(unicode.ToUpper(c))
            } else {
                outStr += string(unicode.ToLower(c))
            }
        } else if unicode.IsSpace(c) { // find a space delimit
            hump = true
        } else {
            if hump { // a space delimit is found before, add a hump
                outStr += string(unicode.ToUpper(c))
            } else {
                outStr += string(unicode.ToLower(c))
            }
            hump = false
        }
    }
    return outStr
}


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    scanner := bufio.NewScanner(os.Stdin)
    var nameStr, outStr string
    for {
        scanner.Scan()
        text := scanner.Text()
        if len(text) == 0 {
            break
        }
        sp := strings.Split(text, ";")
        nameStr = sp[2]
        if sp[0] == "S" {
            if sp[1] == "M" {
                nameStr = strings.Replace(nameStr, "()", "", 1)
            }
            outStr = split(nameStr)
        } else { // sp[0] is "C"
            outStr = combine(sp[1], nameStr)
            if sp[1] == "M" {
                outStr += "()"
            }
        }
        fmt.Println(outStr)
    }
}
