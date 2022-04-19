package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Write your code here
    var outString string = s
    sp := strings.Split(s, ":")
    ch := sp[2][2]
    hour,_ := strconv.Atoi(sp[0])
    if ch == 'P' {
        if hour != 12 {
            hour += 12
            outString = strings.Replace(s, sp[0], strconv.Itoa(hour), 1)
        }
    } else { // tail of s is "AM"
        if hour == 12 {
            outString = strings.Replace(s, sp[0], "00", 1)
        }
    }
    return outString[:8]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

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

