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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var (
        positive float32 = 0
        negative float32 = 0
        zero float32 = 0
        length float32 = float32(len(arr))
    )
    for _, num := range arr {
        //fmt.Printf("%d ", num)
        if num > 0 {
            positive ++
        } else if num < 0 {
            negative ++
        } else {
            zero ++
        }
    }
    //fmt.Printf("p:%f\nn:%f\nz:%f\nl:%f\n", positive, negative, zero, length)
    positive = positive / length
    negative = negative / length
    zero = zero / length
    
    fmt.Printf("%f\n%f\n%f\n", positive, negative, zero)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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

