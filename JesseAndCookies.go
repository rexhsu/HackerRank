package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "container/heap"
)

type Int32Heap []int32

func (h Int32Heap) Len() int { return len(h) }
func (h Int32Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Int32Heap) Swap(i, j int)   { h[i], h[j] = h[j], h[i] }  

func (h *Int32Heap) Push(x interface{}) {
    *h = append(*h, x.(int32))
}

func (h *Int32Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

/*
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func cookies(k int32, A []int32) int32 {
    // Write your code here
    h := &Int32Heap{}
    heap.Init(h)
    
    var (
        least int32
        secLeast int32
        operations int32 = 0
    )
    
    for _ ,num := range A {
        heap.Push(h, num)
    }

    fmt.Println("heap", h, "k", k)
    for h.Len() >= 2 {
        least = heap.Pop(h).(int32)
        if least >= int32(k) {
            return operations
        }
        secLeast = heap.Pop(h).(int32)
        heap.Push(h, least + 2*secLeast)
        operations ++
        //fmt.Println("--- operation", operations, "heap", h)
    }
    
    if h.Pop().(int32) < k {
        return -1
    } else {
        return operations
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var A []int32

    for i := 0; i < int(n); i++ {
        AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
        checkError(err)
        AItem := int32(AItemTemp)
        A = append(A, AItem)
    }

    result := cookies(k, A)

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

